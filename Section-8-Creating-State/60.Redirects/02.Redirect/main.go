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
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8086", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar: ", req.Method)
	// we could process form submission here
	// Earlier we were setting the header and then wrtitng the header.
	// Reason phrase - http.StatusSeeother
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func barred(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred: ", req.Method)
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
