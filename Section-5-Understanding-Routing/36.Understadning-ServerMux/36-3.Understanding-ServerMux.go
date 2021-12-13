package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog doggy")
}

type hotcat int

func (n hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotdog

	// This is to make it a little more elegant.
	// Our net/https package in golang has something called "DefaultServeMux"
	// ListenAndServe start an http server with a given address and handler.
	// The handler is usually 'nil' whihc means that ListenAndServe will use the "DefaultServeMux"
	// So instead of creating our own NewServeMux(), we pass nil to http.ListenAndServe() for the handler.
	// mux := http.NewServeMux()
	// mux.Handle("/dog/", d)
	// mux.Handle("/cat/", c)
	http.Handle("/dog/", d)
	http.Handle("/cat/", c)
	// Handle function is asking ofr a handler. Handler a type that is
	// implementing the Handler interface which means that that type has a method
	// ServeHTTP(ResponseWriter, *Request)
	// so now when we use default servemux, instead of using the handle attach to pointer of serveMux
	// func (m *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request)
	// we use
	// func Handle(pattern string, handler Handler) // This handle funciton is with the net/http package.

	// http.ListenAndServe(":8086", mux)
	http.ListenAndServe(":8086", nil)

}
