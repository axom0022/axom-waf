package main

import (
	"net"
	"net/http"
	"strings"
	"time"
)

func wafhandler(w http.ResponseWriter, r *http.Request) {
	clientip := getrealip(r)
	if iswhitelisted(clientip) {
		reverseproxy.ServeHTTP(w, r)
		return
	}
	if isblacklisted(clientip) {
		http.Error(w, "forbidden", 403)
		return
	}
	recordhttp(clientip)
	if !allowrequest(clientip) {
		addiptablesblock(clientip)
		http.Error(w, "rate limit exceeded", 429)
		return
	}
	if ismalicious(r.URL.Path + " " + r.URL.RawQuery + readbody(r)) {
		addiptablesblock(clientip)
		http.Error(w, "malicious request blocked", 403)
		return
	}
	if isslowloris(r) {
		recordslowloris(clientip)
		http.Error(w, "slowloris blocked", 408)
		return
	}
	reverseproxy.ServeHTTP(w, r)
}

func getrealip(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return strings.Split(xff, ",")[0]
	}
	if xri := r.Header.Get("X-Real-Ip"); xri != "" {
		return xri
	}
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}

func iswhitelisted(ip string) bool {
	for _, w := range config.whitelist {
		if w == ip {
			return true
		}
	}
	return false
}

func isblacklisted(ip string) bool {
	for _, b := range config.blacklist {
		if b == ip {
			return true
		}
	}
	return false
}

func readbody(r *http.Request) string {
	if r.Body == nil {
		return ""
	}
	buf := make([]byte, 1024)
	n, _ := r.Body.Read(buf)
	return string(buf[:n])
}

func isslowloris(r *http.Request) bool {
	if r.Header.Get("Content-Length") == "" && r.ContentLength == 0 && r.Method == "POST" {
		return true
	}
	if r.Header.Get("Transfer-Encoding") == "chunked" {
		return true
	}
	dur := time.Since(r.Context().Done())
	if dur > 10*time.Second {
		return true
	}
	return false
}
