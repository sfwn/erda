package daemonset

import (
	"bytes"
	"encoding/json"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"

	"github.com/erda-project/erda/modules/scheduler/executor/plugins/k8s/k8sapi"
	"github.com/erda-project/erda/modules/scheduler/executor/plugins/k8s/k8serror"
	"github.com/erda-project/erda/pkg/httpclient"
)

type Daemonset struct {
	addr   string
	client *httpclient.HTTPClient
}

type Option func(*Daemonset)

// WithCompleteParams provides an Option
func WithCompleteParams(addr string, client *httpclient.HTTPClient) Option {
	return func(d *Daemonset) {
		d.addr = addr
		d.client = client
	}
}

func New(options ...Option) *Daemonset {
	ds := &Daemonset{}
	for _, op := range options {
		op(ds)
	}
	return ds
}

func (d *Daemonset) Create(ds *appsv1.DaemonSet) error {
	var b bytes.Buffer

	resp, err := d.client.Post(d.addr).
		Path("/apis/apps/v1/namespaces/" + ds.Namespace + "/daemonsets").
		JSONBody(ds).
		Do().
		Body(&b)
	if err != nil {
		return fmt.Errorf("failed to create ds, %s/%s", ds.Namespace, ds.Name)
	}
	if !resp.IsOK() {
		errMsg := fmt.Sprintf("failed to create ds, statuscode: %v, body: %v", resp.StatusCode(), b.String())
		return fmt.Errorf(errMsg)
	}
	return nil
}

func (d *Daemonset) Get(namespace, name string) (*appsv1.DaemonSet, error) {
	var b bytes.Buffer
	resp, err := d.client.Get(d.addr).
		Path("/apis/apps/v1/namespaces/" + namespace + "/daemonsets/" + name).
		Do().
		Body(&b)
	if err != nil {
		return nil, fmt.Errorf("failed to get ds info, %s/%s", namespace, name)
	}
	if !resp.IsOK() {
		if resp.IsNotfound() {
			return nil, k8serror.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get ds info, %s/%s, statuscode: %v, body: %v",
			namespace, name, resp.StatusCode(), b.String())
	}
	ds := &appsv1.DaemonSet{}
	if err := json.NewDecoder(&b).Decode(ds); err != nil {
		return nil, err
	}
	return ds, nil
}

func (d *Daemonset) List(namespace string) (appsv1.DaemonSetList, error) {
	var dsList appsv1.DaemonSetList

	var b bytes.Buffer
	resp, err := d.client.Get(d.addr).
		Path("/apis/apps/v1/namespaces/" + namespace + "/daemonsets").
		Do().
		Body(&b)
	if err != nil {
		return dsList, fmt.Errorf("failed to get ds list, ns: %s, %v", namespace, err)
	}
	if !resp.IsOK() {
		if resp.IsNotfound() {
			return dsList, k8serror.ErrNotFound
		}
		return dsList, fmt.Errorf("failed to get ds list, ns: %s, statuscode: %v, body: %v",
			namespace, resp.StatusCode(), b.String())
	}
	if err := json.NewDecoder(&b).Decode(&dsList); err != nil {
		return dsList, err
	}
	return dsList, nil
}

func (d *Daemonset) Update(daemonset *appsv1.DaemonSet) error {
	var b bytes.Buffer
	resp, err := d.client.Put(d.addr).
		Path("/apis/apps/v1/namespaces/" + daemonset.Namespace + "/daemonsets/" + daemonset.Name).
		JSONBody(daemonset).
		Do().
		Body(&b)
	if err != nil {
		return fmt.Errorf("failed to put daemonset, %s/%s, %v", daemonset.Namespace, daemonset.Name, err)
	}
	if !resp.IsOK() {
		return fmt.Errorf("failed to put daemonset, %s/%s, statuscode: %v, body: %v",
			daemonset.Namespace, daemonset.Name, resp.StatusCode(), b.String())
	}
	return nil
}

func (d *Daemonset) Delete(namespace, name string) error {
	var b bytes.Buffer
	resp, err := d.client.Delete(d.addr).
		Path("/apis/apps/v1/namespaces/" + namespace + "/daemonsets/" + name).
		JSONBody(k8sapi.DeleteOptions).
		Do().
		Body(&b)
	if err != nil {
		return fmt.Errorf("failed to delete daemonsets, %s/%s, %v", namespace, name, err)
	}
	if !resp.IsOK() {
		if resp.IsNotfound() {
			return k8serror.ErrNotFound
		}
		return fmt.Errorf("failed to delete daemonsets, %s/%s, statuscode: %v, body: %v",
			namespace, name, resp.StatusCode(), b.String())
	}
	return nil
}
