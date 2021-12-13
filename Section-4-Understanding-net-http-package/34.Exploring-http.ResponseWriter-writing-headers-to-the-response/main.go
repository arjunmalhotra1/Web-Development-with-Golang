package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// Any type that implements ServHRRP is also of the type Handler.
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Here we are just setting whatever headers we want. "Mcleod-Key" is a header we made up.
	// when we call Header() using a response writer we get back a "Header" type.
	// That header type has a "Set()" method associated with it which allows us to set a custom  header.
	w.Header().Set("Mcleod-Key", "this is from mcleod")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// If we just pass in the Response Writer "w" to Fprintln it writes
	// "<h1>Any code you want in this function</h1>"
	// back as a response.
	fmt.Fprintln(w, "<h1>Any code you want in this function</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8086", d)
}
