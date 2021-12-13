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
	// We are explicitly setting the location to "/" in the response header.
	// And we are setting the WriteHeader to "StatusSeeOther"
	// So the redirect method is always going to be "GET".
	// And the redirect will be sent to "/" this location. Status code will be 303.
	// If we want, we can process the form data here.
	res.Header().Set("Location", "/") // This will redirect to the function "foo".
	// Since we are using http.StatusSeeOther it will change the method to "GET".
	res.WriteHeader(http.StatusSeeOther)
}

func barred(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred: ", req.Method)
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
