(function(){var global=this;(function(){if(!this.require){var e={},t={},n=function o(n,a){var u,s,c=r(a,n),d=r(c,"./index");if(u=t[c]||t[d])return u;if(s=e[c]||e[c=d])return u={id:c,exports:{}},t[c]=u.exports,s(u.exports,function(e){return o(e,i(c))},u),t[c]=u.exports;throw"module "+n+" not found"},r=function(e,t){var n,r,i=[];n=/^\.\.?(\/|$)/.test(t)?[e,t].join("/").split("/"):t.split("/");for(var o=0,a=n.length;o<a;o++)r=n[o],".."==r?i.pop():"."!=r&&""!=r&&i.push(r);return i.join("/")},i=function(e){return e.split("/").slice(0,-1).join("/")};this.require=function(e){return n(e,"")},this.require.define=function(t){for(var n in t)e[n]=t[n]},this.require.modules=e,this.require.cache=t}return this.require}).call(global);var _typeof="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e};window.JSON||(window.JSON={parse:function parse(sJSON){return eval("("+sJSON+")")},stringify:function(){var e=Object.prototype.toString,t=Array.isArray||function(t){return"[object Array]"===e.call(t)},n={'"':'\\"',"\\":"\\\\","\b":"\\b","\f":"\\f","\n":"\\n","\r":"\\r","\t":"\\t"},r=function(e){return n[e]||"\\u"+(e.charCodeAt(0)+65536).toString(16).substr(1)},i=/[\\"\u0000-\u001F\u2028\u2029]/g;return function o(n){if(null==n)return"null";if("number"==typeof n)return isFinite(n)?n.toString():"null";if("boolean"==typeof n)return n.toString();if("object"===("undefined"==typeof n?"undefined":_typeof(n))){if("function"==typeof n.toJSON)return o(n.toJSON());if(t(n)){for(var a="[",u=0;u<n.length;u++)a+=(u?", ":"")+o(n[u]);return a+"]"}if("[object Object]"===e.call(n)){var s=[];for(var c in n)n.hasOwnProperty(c)&&s.push(o(c)+": "+o(n[c]));return"{"+s.join(", ")+"}"}}return'"'+n.toString().replace(i,r)+'"'}}()}),window.performance=window.performance||window.webkitPerformance||window.msPerformance||window.mozPerformance,window.timing=window.timing||{getTimes:function(e){var t=window.performance||window.webkitPerformance||window.msPerformance||window.mozPerformance;if(void 0===t)return!1;var n=t.timing,r={};if(e=e||{},n){if(e&&!e.simple)for(var i in n)!isNaN(parseFloat(n[i]))&&isFinite(n[i])&&(r[i]=parseFloat(n[i]));if(void 0===r.firstPaint){var o=0;if(window.chrome){var a=t.getEntriesByType("paint");if(void 0!==a&&a.length>0){var u={};a.forEach(function(e){u[e.name]=e.startTime}),r.firstPaintTime=u["first-paint"]}else window.chrome.loadTimes&&(o=1e3*window.chrome.loadTimes().firstPaintTime,r.firstPaintTime=o-n.navigationStart)}else"number"==typeof n.msFirstPaint&&(o=n.msFirstPaint,r.firstPaintTime=o-n.navigationStart);e&&!e.simple&&(r.firstPaint=o)}r.loadTime=n.loadEventEnd-n.fetchStart,r.domReadyTime=n.domComplete-n.domInteractive,r.readyStart=n.fetchStart-n.navigationStart,r.redirectTime=n.redirectEnd-n.redirectStart,r.appcacheTime=n.domainLookupStart-n.fetchStart,r.unloadEventTime=n.unloadEventEnd-n.unloadEventStart,r.lookupDomainTime=n.domainLookupEnd-n.domainLookupStart,r.connectTime=n.connectEnd-n.connectStart,r.reponseTime=n.responseEnd-n.responseStart,r.requestTime=n.responseEnd-n.requestStart,r.initDomTreeTime=n.domInteractive-n.responseEnd,r.scriptExecuteTime=n.domContentLoadedEventEnd-n.domContentLoadedEventStart,r.loadEventTime=n.loadEventEnd-n.loadEventStart}return r}};var _typeof="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e};!function(e){if("object"==("undefined"==typeof exports?"undefined":_typeof(exports))&&"undefined"!=typeof module)module.exports=e();else if("function"==typeof define&&define.amd)define([],e);else{var t;t="undefined"!=typeof window?window:"undefined"!=typeof global?global:"undefined"!=typeof self?self:this,t.uuidv1=e()}}(function(){return function e(t,n,r){function i(a,u){if(!n[a]){if(!t[a]){var s="function"==typeof require&&require;if(!u&&s)return s(a,!0);if(o)return o(a,!0);var c=new Error("Cannot find module '"+a+"'");throw c.code="MODULE_NOT_FOUND",c}var d=n[a]={exports:{}};t[a][0].call(d.exports,function(e){var n=t[a][1][e];return i(n?n:e)},d,d.exports,e,t,n,r)}return n[a].exports}for(var o="function"==typeof require&&require,a=0;a<r.length;a++)i(r[a]);return i}({1:[function(e,t,n){function r(e,t){var n=t||0,r=i;return r[e[n++]]+r[e[n++]]+r[e[n++]]+r[e[n++]]+"-"+r[e[n++]]+r[e[n++]]+"-"+r[e[n++]]+r[e[n++]]+"-"+r[e[n++]]+r[e[n++]]+"-"+r[e[n++]]+r[e[n++]]+r[e[n++]]+r[e[n++]]+r[e[n++]]+r[e[n++]]}for(var i=[],o=0;o<256;++o)i[o]=(o+256).toString(16).substr(1);t.exports=r},{}],2:[function(e,t,n){var r="undefined"!=typeof crypto&&crypto.getRandomValues.bind(crypto)||"undefined"!=typeof msCrypto&&msCrypto.getRandomValues.bind(msCrypto);if(r){var i=new Uint8Array(16);t.exports=function(){return r(i),i}}else{var o=new Array(16);t.exports=function(){for(var e,t=0;t<16;t++)0===(3&t)&&(e=4294967296*Math.random()),o[t]=e>>>((3&t)<<3)&255;return o}}},{}],3:[function(e,t,n){function r(e,t,n){var r=t&&n||0,d=t||[];e=e||{};var l=e.node||i,f=void 0!==e.clockseq?e.clockseq:o;if(null==l||null==f){var p=a();null==l&&(l=i=[1|p[0],p[1],p[2],p[3],p[4],p[5]]),null==f&&(f=o=16383&(p[6]<<8|p[7]))}var m=void 0!==e.msecs?e.msecs:(new Date).getTime(),v=void 0!==e.nsecs?e.nsecs:c+1,h=m-s+(v-c)/1e4;if(h<0&&void 0===e.clockseq&&(f=f+1&16383),(h<0||m>s)&&void 0===e.nsecs&&(v=0),v>=1e4)throw new Error("uuid.v1(): Can't create more than 10M uuids/sec");s=m,c=v,o=f,m+=122192928e5;var y=(1e4*(268435455&m)+v)%4294967296;d[r++]=y>>>24&255,d[r++]=y>>>16&255,d[r++]=y>>>8&255,d[r++]=255&y;var b=m/4294967296*1e4&268435455;d[r++]=b>>>8&255,d[r++]=255&b,d[r++]=b>>>24&15|16,d[r++]=b>>>16&255,d[r++]=f>>>8|128,d[r++]=255&f;for(var g=0;g<6;++g)d[r+g]=l[g];return t?t:u(d)}var i,o,a=e("./lib/rng"),u=e("./lib/bytesToUuid"),s=0,c=0;t.exports=r},{"./lib/bytesToUuid":1,"./lib/rng":2}]},{},[3])(3)}),this.require.define({"lib/base-data":function(e,t,n){"use strict";var r=t("./helpers"),i=t("./domain"),o=t("./uuid"),a=t("./"),u=a.browser,s=a.document,c=s.documentPathname(),d=u.userAgent(),l=s.documentDomain(),f={};n.exports={init:function(e){f=e},setUser:function(e){var t=f,n=t.udata;n.uid=e},getOpts:function(){return r.extend({},f)},get:function(){var e=f,t=e.ak,n=e.vid,a=e.url,u=e.udata,s=e.attribute;u.domain=i(f);var p=global.Cookies.get("taid")||o();return r.extend({cid:p,dp:c,dh:l,ua:d,vid:n,url:a,date:Date.now(),attribute:s},u,t?{ak:t}:{})}}}}),this.require.define({"lib/browser/flash-version":function(e,t,n){"use strict";n.exports=function(){var e=0;if(document.all)try{var t=new ActiveXObject("ShockwaveFlash.ShockwaveFlash");if(t){var n=t.GetVariable("$version");e=parseInt(n.split(" ")[1].split(",")[0],10)}}catch(r){}else if(navigator.plugins&&navigator.plugins.length>0){var i=navigator.plugins["Shockwave Flash"];if(i)for(var o=i.description.split(" "),a=0;a<o.length;++a)isNaN(parseInt(o[a],10))||(e=parseInt(o[a],10))}return e}}}),this.require.define({"lib/browser/is-cookie-enabled":function(e,t,n){"use strict";n.exports=function(){return navigator.cookieEnabled?1:0}}}),this.require.define({"lib/browser/screen-color-depth":function(e,t,n){"use strict";n.exports=function(){return window.screen.colorDepth+"-bit"}}}),this.require.define({"lib/browser/screen-size":function(e,t,n){"use strict";n.exports=function(){var e=window,t=e.screen;return t.width+"x"+t.height}}}),this.require.define({"lib/browser/user-agent":function(e,t,n){"use strict";n.exports=function(){return navigator.userAgent}}}),this.require.define({"lib/browser/user-language":function(e,t,n){"use strict";n.exports=function(){return navigator.language||navigator.userLanguage}}}),this.require.define({"lib/browser/window-size":function(e,t,n){"use strict";n.exports=function(){var e=document,t=e.body,n=e.documentElement,r=window.innerWidth||n.clientWidth||t.clientWidth,i=window.innerHeight||n.clientHeight||t.clientHeight;return r+"x"+i}}}),this.require.define({"lib/document/document-character-set":function(e,t,n){"use strict";n.exports=function(){var e=document;return e.characterSet||e.charset||e.inputEncoding}}}),this.require.define({"lib/document/document-domain":function(e,t,n){"use strict";n.exports=function(){return document.domain}}}),this.require.define({"lib/document/document-keywords":function(e,t,n){"use strict";n.exports=function(){for(var e=document.getElementsByTagName("meta"),t="",n=0;n<e.length;n++)if("keywords"===e[n].name){t=e[n].content;break}return t}}}),this.require.define({"lib/document/document-pathname":function(e,t,n){"use strict";n.exports=function(){return window.location.pathname}}}),this.require.define({"lib/document/document-referrer":function(e,t,n){"use strict";n.exports=function(){return document.referrer}}}),this.require.define({"lib/document/document-size":function(e,t,n){"use strict";n.exports=function(){var e=document,t=e.body,n=e.documentElement,r=Math.max(t.scrollWidth,t.offsetWidth,n.clientWidth,n.scrollWidth,n.offsetWidth),i=Math.max(t.scrollHeight,t.offsetHeight,n.clientHeight,n.scrollHeight,n.offsetHeight);return r+"x"+i}}}),this.require.define({"lib/document/document-title":function(e,t,n){"use strict";n.exports=function(){return document.title}}}),this.require.define({"lib/document/document-url":function(e,t,n){"use strict";n.exports=function(){return document.URL}}}),this.require.define({"lib/domain":function(e,t,n){"use strict";n.exports=function(e){return e.udata.domain||window.location.hostname.split(".").slice(-2).join(".")}}}),this.require.define({"lib/events/click":function(e,t,n){"use strict";var r=t("../send-data"),i=r.requestEvent,o=t("../helpers"),a=function(e){var t=document,n=t.documentElement.scrollLeft||t.body.scrollLeft,r=t.documentElement.scrollTop||t.body.scrollTop,a=e.pageX||e.clientX+n,u=e.pageY||e.clientY+r;i({xp:o.getXPath(e.target||e.srcElement).join("/"),x:a,y:u})};n.exports=function(){o.addEventListener(document,"click",function(e){return a(e)})}}}),this.require.define({"lib/events/error":function(e,t,n){"use strict";var r=t("../send-data"),i=r.requestLoadError,o=t("../helpers"),a=function(e){var t=e.target,n="";switch(t.nodeName){case"LINK":n=t.href;break;case"SCRIPT":case"IMG":n=t.src}i({xp:o.getXPath(t||e.srcElement).join("/"),nn:t.nodeName,sr:n})};n.exports=a}}),this.require.define({"lib/events/init":function(e,t,n){"use strict";var r=t("../send-data"),i=r.requestBrowser;n.exports=function(){var e=t("../base-data").get(),n=global.Cookies.get("taid");n||(global.Cookies.set("taid",e.cid,{expires:730,domain:e.domain}),i())}}}),this.require.define({"lib/events/location":function(e,t,n){"use strict";function r(){var e=Event,t=u.version&&u.version>=9&&u.version<=11;if(t){var n=function(e,t){t=t||{bubbles:!1,cancelable:!1,detail:void 0};var n=document.createEvent("CustomEvent");return n.initCustomEvent(e,t.bubbles,t.cancelable,t.detail),n};n.prototype=window.Event.prototype,e=n}var r=function(t){var n=window.history[t];return function(){for(var r=arguments.length,i=Array(r),o=0;o<r;o++)i[o]=arguments[o];var a=n.apply(this,i),u=new e(t);return u.arguments=i,window.dispatchEvent(u),a}};return r}var i=t("../send-data"),o=i.requestDocument,a=t("../performance/spa-timing"),u=t("../ie-version"),s=t("../helpers");n.exports=function(){var e=window.history.pushState;if(e){var t=window.MutationObserver,n=r();window.history.pushState=n("pushState"),window.history.replaceState=n("replaceState");var i=!1,u=window.location,c=u.href,d=function(){o({lc:+(c!==u.href)}),c=u.href,t&&(i=!0,a.onEventStart())};if(s.addEventListener(window,"pushState",d),s.addEventListener(window,"replaceState",d),s.addEventListener(window,"popstate",d),t){var l=window.MutationObserver,f=new l(function(){i&&(a.onEventEnd(),i=!1)});f.observe(document.body,{childList:!0,subtree:!0,characterData:!0})}}}}}),this.require.define({"lib/events/onload":function(e,t,n){"use strict";var r=t("../"),i=r.performance,o=t("../send-data"),a=o.requestDocument,u=o.requestTiming,s=t("../helpers");n.exports=function(e){s.windowOnload(window,function(){setTimeout(function(){var e=i.navigationTiming(),n=i.resourceTiming();u({nt:e,rt:JSON.stringify(n)});var r=t("../base-data").getOpts();a(r.ck?{ck:document.cookie}:void 0)},0)})}}}),this.require.define({"lib/events/unload":function(e,t,n){"use strict";var r=t("../helpers"),i=t("../send-data"),o=i.requestDocument;n.exports=function(){var e=void 0;r.windowOnload(window,function(){setTimeout(function(){e=Date.now()},0)}),r.addEventListener(window,"beforeunload",function(){var t=Math.round((Date.now()-e)/1e3);o({tp:t,lc:0})})}}}),this.require.define({"lib/helpers":function(e,t,n){"use strict";function r(e,t){var n=m.length,r=e.constructor,i=o.isFunction(r)&&r.prototype||u,a="constructor";for(o.has(e,a)&&!o.contains(t,a)&&t.push(a);n--;)a=m[n],a in e&&e[a]!==i[a]&&!o.contains(t,a)&&t.push(a)}var i="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e},o={},a=Object.keys,u=Object.prototype,s=u.hasOwnProperty,c=function(e){return function(t){return null==t?void 0:t[e]}},d=c("length"),l=Math.pow(2,53)-1,f=function(e){var t=d(e);return"number"==typeof t&&t>=0&&t<=l};"function"!=typeof/./&&"object"!==("undefined"==typeof Int8Array?"undefined":i(Int8Array))&&(o.isFunction=function(e){return"function"==typeof e||!1});var p=!{toString:null}.propertyIsEnumerable("toString"),m=["valueOf","isPrototypeOf","toString","propertyIsEnumerable","hasOwnProperty","toLocaleString"],v=function(e,t){return function(){for(var n=arguments.length,r=Array(n),i=0;i<n;i++)r[i]=arguments[i];var o=r.length,a=r[0]||null;if(o<2||null==a)return a;for(var u=1;u<o;u++)for(var s=r[u],c=e(s),d=c.length,l=0;l<d;l++){var f=c[l];t&&void 0!==a[f]||(a[f]=s[f])}return a}};o.allKeys=function(e){if(!o.isObject(e))return[];var t=[];for(var n in e)t.push(n);return p&&r(e,t),t},o.extend=v(o.allKeys),o.isNaN=function(e){return"number"==typeof e&&e!==+e},o.values=function(e){for(var t=o.keys(e),n=t.length,r=Array(n),i=0;i<n;i++)r[i]=e[t[i]];return r},o.contains=o.includes=o.include=function(e,t,n,r){return f(e)||(e=o.values(e)),("number"!=typeof n||r)&&(n=0),o.indexOf(e,t,n)>=0},o.keys=function(e){if(!o.isObject(e))return[];if(a)return a(e);var t=[];for(var n in e)o.has(e,n)&&t.push(n);return p&&r(e,t),t},o.isObject=function(e){var t="undefined"==typeof e?"undefined":i(e);return"function"===t||"object"===t&&!!e},o.has=function(e,t){return null!=e&&s.call(e,t)},o.addEventListener=function(e,t,n,r){document.attachEvent?e.attachEvent("on"+t,n):e.addEventListener(t,n,r)},o.windowOnload=function(e,t){var n=e.document;n.readyState&&"complete"===n.readyState?t():o.addEventListener(e,"load",t)};var h=/^data:.*?\;base64\,.*?==$/i;o.isBase64=function(e){return h.test(e)};var y=function b(e,t){var n=0;if(t=t||[],e.parentNode&&(t=b(e.parentNode,t)),e.previousSibling){n=1;var r=e.previousSibling;do 1===r.nodeType&&r.nodeName===e.nodeName&&n++,r=r.previousSibling;while(r);1===n&&(n=null)}else if(e.nextSibling){var i=e.nextSibling;do 1===i.nodeType&&i.nodeName===e.nodeName?(n=1,i=null):(n=null,i=i.previousSibling);while(i)}if(1===e.nodeType){var o="";e.id&&(o+="[@id='"+e.id+"']"),null!==e.getAttribute("class")&&(o+="[@class='"+e.getAttribute("class")+"']"),o+=n>0?"["+n+"]":"",t.push(e.nodeName.toLowerCase()+o)}return t};o.getXPath=y,n.exports=o}}),this.require.define({"lib/ie-version":function(e,t,n){"use strict";var r=window,i=r.document,o=i.createElement("input"),a=function(){return void 0===r.ActiveXObject?null:r.XMLHttpRequest?i.querySelector?i.addEventListener?r.atob?o.dataset?11:10:9:8:7:6}(),u=!1;if(a&&i.documentMode)try{o.style.behavior="url(#default#clientcaps)",u=i.documentMode!==r.parseInt(o.getComponentVersion("{45EA75A0-A269-11D1-B5BF-0000F8051515}","componentid"))}catch(s){}n.exports={version:a,emulated:!!a&&u}}}),this.require.define({"lib/index":function(e,t,n){"use strict";n.exports={browser:{isCookieEnabled:t("./browser/is-cookie-enabled"),userAgent:t("./browser/user-agent"),userLanguage:t("./browser/user-language"),windowSize:t("./browser/window-size"),screenSize:t("./browser/screen-size"),screenColorDepth:t("./browser/screen-color-depth"),flashVersion:t("./browser/flash-version")},document:{documentSize:t("./document/document-size"),documentTitle:t("./document/document-title"),documentReferrer:t("./document/document-referrer"),documentUrl:t("./document/document-url"),documentCharacterSet:t("./document/document-character-set"),documentPathname:t("./document/document-pathname"),documentDomain:t("./document/document-domain"),documentKeywords:t("./document/document-keywords")},performance:{navigationTiming:t("./performance/navigation-timing"),resourceTiming:t("./performance/resource-timing")}}}}),this.require.define({"lib/log/log-ajax":function(e,t,n){"use strict";var r=t("../send-data"),i=r.requestAjax;n.exports=function(){var e=XMLHttpRequest.prototype.open,t=XMLHttpRequest.prototype.send;XMLHttpRequest.prototype.open=function(){for(var t=arguments.length,n=Array(t),r=0;r<t;r++)n[r]=arguments[r];this.method=n[0].toLowerCase(),this.url=n[1],e.apply(this,n),this.setRequestHeader("terminus-request-id",uuidv1())},XMLHttpRequest.prototype.send=function(){for(var e=this,n=arguments.length,r=Array(n),o=0;o<n;o++)r[o]=arguments[o];this.startTime=Date.now();var a=function(){if(2===e.readyState&&(e.sendTime=Date.now()),3===e.readyState&&(e.loadTime=Date.now()),4===e.readyState){var t=Date.now();i({tt:t-e.startTime,url:e.url,st:e.status,me:e.method,req:JSON.stringify(r||"").length/1024,res:JSON.stringify(e.response||e.reponseText||"").length/1024})}};"function"==typeof window.addEventListener&&this.addEventListener("readystatechange",a),t.apply(this,r)}}}}),this.require.define({"lib/log/log-exec":function(e,t,n){"use strict";var r=t("../send-data"),i=r.requestAjaxError;n.exports=function(e,t,n,r,o){return!("Script error."===e||!t)&&(setTimeout(function(){for(var a=arguments.length,u=Array(a),s=0;s<a;s++)u[s]=arguments[s];r=r||window.event&&window.event.errorCharacter||0;var c="";if(o&&o.stack)c=o.stack.toString();else if(u.callee){for(var d=[],l=u.callee.caller,f=3;l&&--f>0&&(d.push(l.toString()),l!==l.caller);)l=l.caller;d=d.join(","),c=d}i({ers:t,erm:e,erl:n,erc:r,sta:c})},0),!1)}}}),this.require.define({"lib/performance/navigation-timing":function(e,t,n){"use strict";var r=t("./utility");n.exports=function(){return r.compressNavigationTiming(window.timing.getTimes())}}}),this.require.define({"lib/performance/resource-timing":function(e,t,n){"use strict";var r=t("./utility");n.exports=function(){return window.performance&&window.performance.getEntriesByType?r.compressResourceTiming(window.performance.getEntriesByType("resource")):{}}}}),this.require.define({"lib/performance/spa-timing":function(e,t,n){"use strict";var r=t("./utility"),i=t("../send-data"),o=i.requestTiming,a={};n.exports={onEventStart:function(){a.startTime=Date.now()},onEventEnd:function(){a.endTime=Date.now();var e=a.endTime-a.startTime;o({nt:r.compressNavigationTiming({loadTime:e})})}}}}),this.require.define({"lib/performance/utility":function(e,t,n){"use strict";var r="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e},i=t("../helpers"),o=function(e){var t={},n=void 0,r=void 0,o=void 0,a=void 0,u=void 0;return i.keys(e).forEach(function(i){if(e.hasOwnProperty(i)&&!(i.length>150)){n=e[i],r=i.split(""),a=t;for(var s=0;s<r.length;s++)o=r[s],u=a[o],"undefined"==typeof u?a=a[o]=s===r.length-1?n:{}:"string"==typeof u?a=a[o]={"|":u}:s===r.length-1?a[o]["|"]=n:a=a[o]}}),t},a=function p(e,t){var n=0,o=void 0,a=void 0,u=void 0,s=[];i.keys(e).forEach(function(t){e.hasOwnProperty(t)&&s.push(t)});for(var c=0;c<s.length;c++)o=s[c],"object"===r(e[o])&&(a=p(e[o],!1),a&&(delete e[o],o+=a.name,e[o]=a.value)),n++;return 1===n?t?(u={},u[o]=e[o],u):{name:o,value:e[o]}:!!t&&e},u=function(e){return"number"==typeof e&&0!==e?e.toString(36):"string"==typeof e?e:""},s=function(e){for(var t=[],n=0;n<e.length;n++)t.push(u(e[n]));return t},c={other:0,img:1,link:2,script:3,css:4,xmlhttprequest:5,iframe:6,image:7},d=function(e,t){"number"!=typeof e&&(e=0),"number"!=typeof t&&(t=0);var n=Math.round(e),r=Math.round(t);return 0===n?0:n-r},l=function(e){for(var t={},n=0;n<e.length;n++){var r=e[n],u=r.name,l=r.initiatorType,f=r.startTime,p=r.responseEnd,m=r.responseStart,v=r.requestStart,h=r.connectEnd,y=r.secureConnectionStart,b=r.connectStart,g=r.domainLookupEnd,w=r.domainLookupStart,x=r.redirectEnd,S=r.redirectStart;if(!i.isBase64(u)&&!t[u]){var E=c[l]||0,q=E+s([d(f),d(p,f),d(m,f),d(v,f),d(h,f),d(y,f),d(b,f),d(g,f),d(w,f),d(x,f),d(S,f)]).join(",").replace(/,+$/,"");t[u]=q}}return t=a(o(t),!0)},f=function(e){var t=e.loadTime,n=e.readyStart,r=e.domReadyTime,i=e.scriptExecuteTime,o=e.requestTime,a=e.reponseTime,u=e.initDomTreeTime,c=e.loadEventTime,d=e.unloadEventTime,l=e.appcacheTime,f=e.connectTime,p=e.lookupDomainTime,m=e.redirectTime;return s([t,n,r,i,o,a,u,c,d,l,f,p,m]).join(",").replace(/,+$/,"")};n.exports={compressResourceTiming:l,compressNavigationTiming:f}}}),this.require.define({"lib/polyfill":function(e,t,n){"use strict";n.exports=function(){Array.prototype.forEach||(Array.prototype.forEach=function(e){var t=this.length;if("function"!=typeof e)throw new TypeError;for(var n=arguments[1],r=0;r<t;r++)r in this&&e.call(n,this[r],r,this)})}}}),this.require.define({"lib/request":function(e,t,n){"use strict";function r(e,t){var n={};for(var r in e)t.indexOf(r)>=0||Object.prototype.hasOwnProperty.call(e,r)&&(n[r]=e[r]);return n}var i="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e},o=t("./helpers"),a=t("./ie-version");"undefined"==typeof window.XMLHttpRequest&&(window.XMLHttpRequest=function(){for(var e=["Microsoft","msxm3","msxml2","msxml1"],t=0;t<e.length;t++)try{var n=e[t]+".XMLHTTP";return new ActiveXObject(n)}catch(r){}});var u=XMLHttpRequest.prototype.open,s=XMLHttpRequest.prototype.send,c=encodeURIComponent,d=function(e){var t=[];return o.keys(e).forEach(function(n){t.push(c(n)+"="+c(e[n]))}),t.join("&")},l=function(e,t){var n=document,r=n.createElement("iframe"),i="terminus-analytics-"+Date.now();n.body.appendChild(r),r.style.display="none",r.contentWindow.name=i;var a=n.createElement("form");a.target=i,a.style.display="none",a.action=e,a.method="POST",o.keys(t).forEach(function(e){var r=n.createElement("input");r.type="hidden",r.name=e,r.value=t[e],a.appendChild(r)}),n.body.appendChild(a),a.submit(),o.addEventListener(r,"load",function(){n.body.removeChild(a),n.body.removeChild(r)})},f=function(e,t,n){if(navigator.sendBeacon&&Blob){var r=new Blob([t],{type:n});navigator.sendBeacon(e,r)}else{var i=new XMLHttpRequest;u.call(i,"POST",e,!0),i.setRequestHeader("Content-type",n),s.call(i,t)}},p=function(e){var n=t("./base-data").get(),r=n.ak,i=n.cid,o=n.uid,u=n.url,s=n.attribute,c={data:e,cid:i};void 0!==r&&(c.ak=r),void 0!==o&&(c.uid=o),void 0!==s&&(c.attribute=s),a.version&&a.version<10?l(u,c):f(u,d(c),"application/x-www-form-urlencoded; charset=UTF-8")};n.exports={requestCustom:function(e){var t=e.data,n=r(e,["data"]),a=d(n).split("&");o.keys(t).forEach(function(e){var n=t[e];o.isObject(n)&&o.keys(n).forEach(function(t){["string","number","boolean"].indexOf(i(n[t]))>-1&&a.push(e+"="+c(t)+"="+c(n[t]))})}),p(a.join("&"))},requestData:function(e){var n=t("./base-data").get(),i=(n.ak,n.cid,n.uid,n.url,n.attribute,r(n,["ak","cid","uid","url","attribute"]));p(d(o.extend(e,i)))}}}}),this.require.define({"lib/send-data":function(e,t,n){"use strict";var r=Object.assign||function(e){for(var t=1;t<arguments.length;t++){var n=arguments[t];for(var r in n)Object.prototype.hasOwnProperty.call(n,r)&&(e[r]=n[r])}return e},i=(t("./helpers"),t("./")),o=i.browser,a=i.document,u=t("./request"),s=u.requestData,c=u.requestCustom;n.exports={requestDocument:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};s(r({t:"document",dt:a.documentTitle(),ds:a.documentSize(),dr:a.documentReferrer(),dl:a.documentUrl(),de:a.documentCharacterSet(),dp:a.documentPathname(),dh:a.documentDomain(),dk:a.documentKeywords(),lc:1},e))},requestBrowser:function(){s({t:"browser",ce:o.isCookieEnabled(),vp:o.windowSize(),ua:o.userAgent(),ul:o.userLanguage(),sr:o.screenSize(),sd:o.screenColorDepth(),fl:o.flashVersion()})},requestTiming:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};s(r({t:"timing"},e))},requestEvent:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};s(r({t:"event"},e))},requestLoadError:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};s(r({t:"loadError"},e))},requestAjax:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};s(r({t:"request"},e))},requestAjaxError:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};s(r({t:"error"},e))},requestMetric:function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};c({t:"metric",n:e,date:Date.now(),data:t})}}}}),this.require.define({"lib/uuid":function(e,t,n){"use strict";n.exports=function(){return"xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g,function(e){var t=16*Math.random()|0,n="x"===e?t:3&t|8;return n.toString(16)})}}}),function(e){global.Cookies=e()}(function(){function e(){for(var e=0,t={};e<arguments.length;e++){var n=arguments[e];for(var r in n)t[r]=n[r]}return t}function t(n){function r(t,i,o){var a;if("undefined"!=typeof document){if(arguments.length>1){if(o=e({path:"/"},r.defaults,o),"number"==typeof o.expires){var u=new Date;u.setMilliseconds(u.getMilliseconds()+864e5*o.expires),o.expires=u}try{a=JSON.stringify(i),/^[\{\[]/.test(a)&&(i=a)}catch(s){}return i=n.write?n.write(i,t):encodeURIComponent(String(i)).replace(/%(23|24|26|2B|3A|3C|3E|3D|2F|3F|40|5B|5D|5E|60|7B|7D|7C)/g,decodeURIComponent),t=encodeURIComponent(String(t)),t=t.replace(/%(23|24|26|2B|5E|60|7C)/g,decodeURIComponent),t=t.replace(/[\(\)]/g,escape),document.cookie=[t,"=",i,o.expires?"; expires="+o.expires.toUTCString():"",o.path?"; path="+o.path:"",o.domain?"; domain="+o.domain:"",o.secure?"; secure":""].join("")}t||(a={});for(var c=document.cookie?document.cookie.split("; "):[],d=/(%[0-9A-Z]{2})+/g,l=0;l<c.length;l++){var f=c[l].split("="),p=f.slice(1).join("=");'"'===p.charAt(0)&&(p=p.slice(1,-1));try{var m=f[0].replace(d,decodeURIComponent);if(p=n.read?n.read(p,m):n(p,m)||p.replace(d,decodeURIComponent),this.json)try{p=JSON.parse(p)}catch(s){}if(t===m){a=p;break}t||(a[m]=p)}catch(s){}}return a}}return r.set=r,r.get=function(e){return r.call(r,e)},r.getJSON=function(){return r.apply({json:!0},[].slice.call(arguments))},r.defaults={},r.remove=function(t,n){r(t,"",e(n,{expires:-1}))},r.withConverter=t,r}return t(function(){})}),window.MutationObserver=window.MutationObserver||window.WebKitMutationObserver||window.MozMutationObserver,this.require.define({ta:function(e,t,n){"use strict";var r=t("./lib/helpers"),i=t("./lib/uuid"),o=t("./lib/ie-version");n.exports=function(e){var n={policy:"blacklist",path:[],vid:i(),url:"//analytics.terminus.io/collect",udata:{},ck:!1};o.version&&o.version<8||(t("./lib/polyfill")(),e=r.extend(n,e),t("./lib/base-data").init(e),t("./lib/events/init")(e),t("./lib/events/unload")(e),t("./lib/events/click")(e),t("./lib/log/log-ajax")(e),t("./lib/events/onload")(e),t("./lib/events/location")(e))}}}),function(){var e=global.require,t=function(e){return Array.prototype.slice.call(e)},n={start:function(){e("ta").apply(null,arguments)},setUser:function(){e("lib/base-data").setUser.apply(null,arguments)},sendError:function(){e("lib/events/error").apply(null,arguments)},sendExecError:function(){e("lib/log/log-exec").apply(null,arguments)},send:function(){e("lib/send-data").requestMetric.apply(null,arguments)}},r=$ta.q,i=function(e){var r=t(e),i=r.shift();n[i].apply(n,r)},o=function(){r.forEach(i),r.length=0,r.push=i};r.push=function(e){"start"===e[0]?(r.unshift(e),o()):Array.prototype.push.call(r,e)};for(var a=-1,u=0;u<r.length;u++)if("start"===r[u][0]){a=u;break}if(a>-1){var s=r.splice(a,1)[0];r.push(s),o()}}.call()}).call({});