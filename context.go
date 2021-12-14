package main

import "net/http"

type Context struct {
	Params         map[string]interface{}
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

type HandlerFunc func(*Context)
