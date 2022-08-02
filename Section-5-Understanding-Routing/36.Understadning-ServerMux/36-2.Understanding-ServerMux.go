// Mux runs a code based on different routes.
// For a certain path and certain Method run a seperate piece of code.
// Old school method was - For a particular path run a different PHP file.
// Paths don't point to a file anymore.
// https://github.com/GoesToEleven/golang-web-dev/tree/master/018_understanding-net-http-ServeMux
// When we have a pointer to servemux we are also implementing the handler interface.
// https://pkg.go.dev/go.reizu.org/servemux
// type ServeMux
// func (m *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request)
// To get a pointer to servemux we use -
// func NewServeMux() *ServeMux
// So we have a value of type pointer to Servemux we can pass that value in where ever handler is needed.
// Handle is a method available to any value that is a pointer to servemux.

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

	mux := http.NewServeMux() // Returns pointer to ServeMux
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)
	// Note - We have "/dog/"
	// because even if something comes like /dog/something/else it will still run the dog handler code.
	// Hence the last trailing '/'
	// path: /cat/something/else will return a 404.
	// Because we don't have a '/' after cat on line 40, the down line paths are not caught.

	// Notice we pass in our mux as a handle since our mux has ServeHTTP(ResponseWriter, *Request)
	http.ListenAndServe(":8086", mux)

}
