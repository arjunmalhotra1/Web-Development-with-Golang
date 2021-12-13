package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8086", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	//body
	// To read the body we make a slice of byted which is the length of the req.ContentLength
	// when we give make only one value it will set both length and capacity to that one value.
	bs := make([]byte, req.ContentLength)
	// Next we read all the req body into that slice of bytes, just using the read method.
	req.Body.Read(bs)
	// Next we convert the byte slice and convert it into string.
	body := string(bs)

	err := tpl.ExecuteTemplate(res, "index.gohtml", body)
	if err != nil {
		http.Error(res, err.Error(), 500)
		log.Fatalln(err)
	}
}
