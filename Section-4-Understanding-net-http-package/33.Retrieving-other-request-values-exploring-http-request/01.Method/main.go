package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // Off of my request parse my form. To get "req.Form" on line 18 we need to first call
	// req.ParseForm() as we do on line 13. This parses our form and makes all the data available to us.
	if err != nil {
		log.Fatalln(err) // log the error and shut down the program in case of an error.
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data) // There is a field attached to the req(req) "Form".
	// http.Request req is a struct in http packge. Here we are jsut saying "Give me my form".
	// Form in req struct is of type "url.Values" which is a map. This map takes a string and then returns a
	// slice of string.
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8086", d) // Anything that comes in from port 8080 that will be handled by the
	//ServerHTTP attached to "d" which is hotdog.
}
