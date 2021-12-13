// Routing - If people make request to our server with diffferent URLs/different paths.
// Routing is how do we respond to different URLs and serve different codes for those URLs.

// ALL THESE WORDS ARE SYNONYMOUS.
//Router
//Request Router
// Multiplexer
// Mux
// ServeMux
// Server
// Http Router
// Http Request Router
// Http Multiplexer
// Http Mux
// Http ServeMux
// Http Sever

package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// This is basic routing nad is not an elegant way of doing routing.
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy doggy doggy")
	case "/cat":
		io.WriteString(res, "kitty kitty kitty")
	}
}

func main() {
	var d hotdog // d is set to 0 value.
	// d is type hotdog, it is also implementing the handler interface and is of underlying type int.
	http.ListenAndServe(":8086", d)
}
