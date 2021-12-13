// This code is same as 36-4.Understanding-ServeMux.go
package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog doggy")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	http.HandleFunc("/dog/", d) // for the route "/dog/" use function d
	http.HandleFunc("/cat", c)  // for the route "/cat" use function c

	http.ListenAndServe(":8086", nil)

}
