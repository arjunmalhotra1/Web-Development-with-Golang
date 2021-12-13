package main

import (
	"io"
	"net/http"
)

//type hotdog int

//func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog doggy")
}

//type hotcat int

// func (n hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	// var d hotdog
	// var c hotdog
	// We can make this better by using Handlefunc instead of http.Handle
	// HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	// https://pkg.go.dev/net/http#HandleFunc
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	// It takes in a pattern and a (identifier) "handler" which is of this type that is a
	// function of type/signature:
	// func(ResponseWriter, *Request)
	// This func(ResponseWriter, *Request) is not a handler type, it is it's own type.
	// Notice the different between Handle and Handlefunc
	// func Handle(pattern string, handler Handler)

	// So now we can get rid of line 21 and 22.
	// Also we can get rid of lines 8 and 15.
	// Also notice the change in the functions.
	// We no longer need
	// func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// because Handlefunc wants a function with this signature - func(ResponseWriter, *Request)
	// http.HandleFunc("/dog/", http.HandlerFunc(d))
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)
	// HandleFucn at the net/http level is attaching itself with the default servemux.

	http.ListenAndServe(":8086", nil)

}
