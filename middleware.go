package main

import (
	"fmt"
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

// POST에 전송된 Form 데이터를 Context의 Param에 담는 미들웨어
func parseFormHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		c.Request.ParseForm()
		fmt.Println(c.Request.PostForm)
		for k, v := range c.Request.PostForm {
			if len(v) > 0 {
				c.Params[k] = v[0]
			}
		}
		next(c)
	}
}
