package main

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]map[string]HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h HandlerFunc) {
	// 라우터 내부에 해당 url에 존재하는 핸들러가 없으면 정의
	m, ok := r.handlers[method]
	if !ok {
		m = make(map[string]HandlerFunc)
		r.handlers[method] = m
	}
	m[pattern] = h
}

func match(pattern, path string) (bool, map[string]string) {
	if pattern == path {
		return true, nil
	}

	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	if len(patterns) != len(paths) {
		return false, nil
	}

	params := make(map[string]string)

	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:

		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			params[patterns[i][1:]] = paths[i]
		default:
			return false, nil
		}
	}

	return true, params
}

// func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	// 핸들러 내부에 정의된 URL 패턴과 요청 URL을 비교해 일치하는 것을 찾는다.
// 	for pattern, handler := range r.handlers[req.Method] {
// 		if ok, params := match(pattern, req.URL.Path); ok {
// 			c := Context{
// 				Params:         make(map[string]interface{}),
// 				ResponseWriter: w,
// 				Request:        req,
// 			}
// 			for k, v := range params {
// 				c.Params[k] = v
// 			}

// 			handler(&c)
// 			return
// 		}
// 	}
// 	http.NotFound(w, req)
// 	return
// }

func (r *router) handler() HandlerFunc {
	return func(c *Context) {
		fmt.Println("무야호")
		for pattern, handler := range r.handlers[c.Request.Method] {

			if ok, params := match(pattern, c.Request.URL.Path); ok {
				fmt.Println("무야호")
				for k, v := range params {
					c.Params[k] = v
				}
				handler(c)
				return
			}
		}

		http.NotFound(c.ResponseWriter, c.Request)
		return
	}
}
