package main

import (
	"fmt"
	"net/http"
)

func main() {

	r := &router{make(map[string]map[string]HandlerFunc)}

	r.HandleFunc("GET", "/", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "welcome!")
	})

	r.HandleFunc("GET", "/about", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "About!")
	})

	r.HandleFunc("GET", "/users/:id", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "retrieve user!", c.Params["id"])
	})

	r.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "retrieve user's address!", c.Params["user_id"], c.Params["id"])
	})

	r.HandleFunc("POST", "/users", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "create user")
	})

	r.HandleFunc("POST", "/users/:user_id/addresses", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "create user's address")
	})

	http.ListenAndServe(":8000", r)
}
