// https://github.com/GoesToEleven/golang-web-dev/tree/master/027_passing-data#passing-values

// When we use the method POST on the form the data goes through the body in the http request(form is the client).
// If we use method GET on the form data goes through the URL.

// We can always send the values/data in the URL regardless of whether it's from form submission.
// If we are talking about state and we need to rememeber who this user is, and they don't allow any cookies.
// Dropping in url values is one of the ways to track people if they have not accepted any cookies.
// We can always drop a unique id into every link on our web page/the URL.
// That universally unique id, is some id that for each client we dynamically generate the page an inser that unique
// id into every link. We can just check that unique id and know which user is the request coming from.
// On any link the user clicks that unique id is coming back to the server.

//func (r *Request) FormValue(key string) string

package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8086", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q") // foo is going ot ask for the request form value by asking for the identifier "q".
	// If q is not present/ or if the key is not present, we get an empty string.
	// If it is present it will get back the value.
	io.WriteString(res, "Do my search: "+v) // so v will always be atleast an empty string.

}

// This is oputting a value into the url by the client - like this - "?q=dog"
// This is the way we pull out that value and display it
// v := req.FormValue("q")
// Using FormValue to retrieve value from the URL.
// If that value as here "q" would have been coming ofrom the form as well the
// POST and PUT takes precedence over the URL query string.
// Visit this page
// http://localhost:8086/?q=dog
