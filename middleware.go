package main

import (
	"log"
	"time"
)

type Middleware func(next HandlerFunc) HandlerFunc

func logHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		// 다음 핸들러를 수행
		next(c)

		log.Printf(
			"[%s] %q %v\n",
			c.Request.Method,
			c.Request.URL.String(),
			time.Since(t))
	}

}
