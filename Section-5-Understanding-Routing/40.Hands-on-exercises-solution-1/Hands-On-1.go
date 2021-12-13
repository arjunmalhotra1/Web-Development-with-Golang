// ListenAndServe on port ":8080" using the default ServeMux.

// Use HandleFunc to add the following routes to the default ServeMux:

// "/" "/dog/" "/me/

// Add a func for each of the routes.

// Have the "/me/" route print out your name.
package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "In the / route")
}
func e(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "In the route /dog/")
}
func f(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "My name is Arjun")
}
func main() {
	http.Handle("/", http.HandlerFunc(d))
	http.Handle("/dog/", http.HandlerFunc(e))
	http.Handle("/me/", http.HandlerFunc(f))
	http.ListenAndServe(":8086", nil)
}
