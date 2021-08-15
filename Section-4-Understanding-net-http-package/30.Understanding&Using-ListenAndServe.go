package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// Underlying type of hotdog is an int
// We are attaching a method to hotdog ServeHTTP.
// Any value of type hotdog is now also a value of type Handler implicitly.
// Hotdog implicitly implements the handler interface.
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d hotdog
	// Anything that ocmes in the port 8080 will be handled by "d"
	http.ListenAndServe(":8086", d) // ListenAndServe is jsut a basic version of what we saw with TCP in Section-3

}
