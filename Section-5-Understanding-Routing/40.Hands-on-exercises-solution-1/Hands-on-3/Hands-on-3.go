// 1. Take the previous program and change it so that:
// * func main uses http.Handle instead of http.HandleFunc

// Contstraint: Do not change anything outside of func main

// Hints:

// [http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)
// ``` Go
// type HandlerFunc func(ResponseWriter, *Request)
// ```

// [http.HandleFunc](https://godoc.org/net/http#HandleFunc)
// ``` Go
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
// ```

// [source code for HandleFunc](https://golang.org/src/net/http/server.go#L2069)
// ``` Go
//   func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
//   		mux.Handle(pattern, HandlerFunc(handler))
//   }
// ```

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
