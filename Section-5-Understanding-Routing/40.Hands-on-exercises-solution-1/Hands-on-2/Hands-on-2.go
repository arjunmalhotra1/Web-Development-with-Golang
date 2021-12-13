// Take the previous program in the previous folder and change it so that:
// a template is parsed and served
// you pass data into the template

package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "In the / route")
}
func e(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "In the route /dog/")
}
func f(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing tempate", err)
	}
	err = tpl.ExecuteTemplate(res, "something.gohtml", "My name is Arjun")
	if err != nil {
		log.Fatalln("Error executing the template", err)
	}
}
func main() {
	http.Handle("/", http.HandlerFunc(d))
	http.Handle("/dog/", http.HandlerFunc(e))
	http.Handle("/me/", http.HandlerFunc(f))
	http.ListenAndServe(":8086", nil)
}
