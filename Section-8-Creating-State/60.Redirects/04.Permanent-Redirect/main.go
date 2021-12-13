package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// We go to barred, "http:localhost:8086/barred"
// barred is going to execute a template with index.gohtml.

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8086", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar: ", req.Method)
	// we could process form submisison here.
	// we will go to "bar" and "bar" will redirect us to "foo"
	// After the first request when we try to go to "bar" again we will not even hit "bar"
	// Browser/Client rememebers that it has been permanently moved.
	// So the browser doesn't send to bar but sends us to "foo", when we hit "localhost:8086/bar"
	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}

// Browsers remembers the Moved Permanently.
// Browsers will not even ask to move to get to that old location anymore. We need to clear our browser cache.
// 301 - Moved Permanently
// 303 - See Other - Always going to be GET in the redirect.
// 307 - Temporary Redirect - Will preserve the method.
// We use 307 when we have a POST method and we want to preserve the POST call in the redirect as well.
