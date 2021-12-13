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
	v := req.FormValue("q")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<form method="get">
		<input type="text" name="q">
		<input type="submit">
		</form>
		<br>`+v)
}

// When the method is get on form submission the query changes
// from
// http://localhost:8086
// to
// http://localhost:8086/?q=70115
// as the q is now passed through the URL.
