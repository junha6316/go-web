package main

import "net/http"

type router struct {
	handlers map[string]map[string]http.HandlerFunc
}

// find handleFunction using pattern
func (r *router) HandleFunc(method, pattern string, h http.HandlerFunc) {

	m, ok := r.handlers[method]
	if !ok {
		m = make(map[string]http.HandlerFunc)
		r.handlers[method] = m
	}
	m[pattern] = h
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if m, ok := r.handlers[req.Method]; ok {
		if h, ok := m[req.URL.Path]; ok {
			h(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
