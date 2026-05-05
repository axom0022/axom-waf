package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

var reverseproxy *httputil.ReverseProxy

func init() {
	target, _ := url.Parse(config.upstreamtarget)
	reverseproxy = httputil.NewSingleHostReverseProxy(target)
}
