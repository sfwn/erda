// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transports

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/http/httpproxy"
	"golang.org/x/net/proxy"
)

const (
	envKeyForwardProxyHosts = "FORWARD_PROXY_HOSTS"
	envKeyForwardHttpProxy  = "FORWARD_HTTP_PROXY"
	envKeyForwardHttpsProxy = "FORWARD_HTTPS_PROXY"
	envKeyNoProxy           = "NO_PROXY"
)

// shared base dialer for all outbound connections
var baseDialer = &net.Dialer{
	Timeout:   10 * time.Second, // 建立 TCP 连接的超时时间，超过则视为连接失败（包括走代理的场景）
	KeepAlive: 60 * time.Second, // TCP keepalive 间隔，用于长连接保活与连接复用
}

// forwardProxyHosts caches the FORWARD_PROXY_HOSTS list for both HTTP and SOCKS proxy decisions.
var forwardProxyHosts = func() []string {
	hosts := strings.Split(os.Getenv(envKeyForwardProxyHosts), ",")
	var cleaned []string
	for _, h := range hosts {
		h = strings.TrimSpace(h)
		if h == "" {
			continue
		}
		cleaned = append(cleaned, h)
	}
	return cleaned
}()

// socksProxyURL and socksDialer are non-nil only when FORWARD_HTTP_PROXY / FORWARD_HTTPS_PROXY
// is configured with a socks5:// scheme.
var (
	socksProxyURL *url.URL
	socksDialer   proxy.Dialer
)

// init forward proxy configuration (HTTP/HTTPS) and optional SOCKS5 dialer.
func init() {
	httpProxy := os.Getenv(envKeyForwardHttpProxy)
	httpsProxy := os.Getenv(envKeyForwardHttpsProxy)

	// If both are set and not equal, fail fast.
	if httpProxy != "" && httpsProxy != "" && httpProxy != httpsProxy {
		panic(fmt.Sprintf("FORWARD_HTTP_PROXY (%q) and FORWARD_HTTPS_PROXY (%q) must be equal when both are set", httpProxy, httpsProxy))
	}

	// Normalize: if only one is set, copy it to the other so that the effective config is consistent.
	if httpProxy == "" {
		httpProxy = httpsProxy
	}
	if httpsProxy == "" {
		httpsProxy = httpProxy
	}

	ProxyConfig = &httpproxy.Config{
		HTTPProxy:  httpProxy,
		HTTPSProxy: httpsProxy,
		NoProxy:    os.Getenv(envKeyNoProxy),
		CGI:        os.Getenv("REQUEST_METHOD") != "",
	}

	// Initialize SOCKS5 dialer if the effective proxy is socks5://
	raw := httpsProxy
	if raw == "" {
		raw = httpProxy
	}
	if raw != "" {
		if u, err := url.Parse(raw); err == nil && strings.EqualFold(u.Scheme, "socks5") {
			auth := &proxy.Auth{}
			if u.User != nil {
				auth.User = u.User.Username()
				if pw, ok := u.User.Password(); ok {
					auth.Password = pw
				}
			}
			if d, err := proxy.SOCKS5("tcp", u.Host, auth, baseDialer); err == nil {
				socksProxyURL = u
				socksDialer = d
			} else {
				logrus.WithError(err).Errorf("failed to init socks5 proxy for %s", raw)
			}
		}
	}

	// Print final proxy configuration at startup.
	socksEnabled := socksProxyEnabled()
	logrus.Infof(
		"ai-proxy forward proxy config: HTTP=%q HTTPS=%q NO_PROXY=%q SOCKS5_ENABLED=%t SOCKS5_URL=%v HOSTS=%v",
		ProxyConfig.HTTPProxy,
		ProxyConfig.HTTPSProxy,
		ProxyConfig.NoProxy,
		socksEnabled,
		socksProxyURL,
		forwardProxyHosts,
	)
}

func socksProxyEnabled() bool {
	return socksDialer != nil && socksProxyURL != nil
}

// ProxyConfig is forward proxy configuration, i.e., proxy configuration for transport outbound traffic
var ProxyConfig = &httpproxy.Config{}

// BaseTransport returns a basic http.RoundTripper. It checks whether the host requested by *http.Request is in the FORWARD_PROXY_HOSTS list,
// if it is in the list, it uses ProxyConfig's proxy configuration, if not in the list, it uses the default proxy configuration http.ProxyFromEnvironment.
var BaseTransport http.RoundTripper = &http.Transport{
	Proxy: func(req *http.Request) (*url.URL, error) {
		// Only use HTTP/HTTPS proxy here. If a SOCKS5 proxy is configured, we rely on DialContext instead.
		for _, host := range forwardProxyHosts {
			if strings.HasSuffix(req.Host, host) || strings.HasSuffix(req.URL.Host, host) {
				if socksProxyEnabled() {
					// For SOCKS, do not configure an HTTP proxy; DialContext will route via SOCKS5.
					return nil, nil
				}
				proxyURL, err := ProxyConfig.ProxyFunc()(req.URL)
				if err != nil {
					return nil, err
				}
				logrus.Debugf("%s use http proxy, proxyURL: %s", host, proxyURL)
				return proxyURL, nil
			}
		}
		return http.ProxyFromEnvironment(req)
	},
	DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
		// When SOCKS5 is enabled and the destination host matches FORWARD_PROXY_HOSTS,
		// route the connection through the SOCKS5 proxy.
		if socksProxyEnabled() {
			host, _, err := net.SplitHostPort(addr)
			if err != nil {
				host = addr
			}
			for _, h := range forwardProxyHosts {
				if strings.HasSuffix(host, h) {
					logrus.Debugf("%s use socks5 dialer via %s", host, socksProxyURL)
					// x/net/proxy.Dialer does not have a context-aware API;
					// we ignore ctx here and rely on underlying dialer timeouts.
					return socksDialer.Dial(network, addr)
				}
			}
		}
		return baseDialer.DialContext(ctx, network, addr)
	},
	TLSHandshakeTimeout:   10 * time.Second,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	ForceAttemptHTTP2:     true,
	DisableCompression:    false,
}
