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
