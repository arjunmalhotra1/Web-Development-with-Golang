// Passing values through form submission.
// When somebody submits a form the values can be sent tot he server in 2 ways.
// In one way it can go is in the request body of the http request.
// Other way we can get values from the users submission of the form is
// by sending those values through the URL.

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
	v := req.FormValue("q") // If no identifier "q" then FormValue will return an empty string.
	// Either way we will haev a value or an empty string.
	// We set our header for our response as HTML.
	// Whatever value for "q" came from form or the url will be printed.
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Since in this form we use method as "post". Since we use post it will
	// send the data through the form's request body.
	io.WriteString(res, `
	<form method="POST">
		<input type="text" name="q">
		<input type="submit">
		</form>
		<br>`+v)
}
