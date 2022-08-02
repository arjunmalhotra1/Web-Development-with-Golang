// Disambiguate between a function with parameters "ResponseWriter and a pointer to Request" vs
// the type handler func
// i.e.
// Differences between '=
// func(ResponseWriter, *Request) and Handlerfunc

// https://pkg.go.dev/net/http#HandlerFunc
// type Handler func is a fucntion of responseWriter and a pointer to request
// type HandlerFunc func(ResponseWriter, *Request)
// Note on 36-4. We use http.handleFunc

// when we have
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
// handler func(ResponseWriter, *Request) this is different from the
// type HandlerFunc(on line 7). HandlerFunc's underlying type is func(ResponseWriter, *Request)

// type hotdog int - type hotdog's underlying type is an int
// type HandlerFunc func(ResponseWriter, *Request) - type handlerFunc's underlying type is a
// func(ResponseWriter, *Request)
// So "HandlerFunc" is it's own type.
// and "HandlerFunc" has a method ServeHTTP with ResponseWriter and pointer to a Request attached to it.
// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
// So "HandlerFunc" can be used in situations where we have a fucntion of type
// func(ResponseWriter, *Request) and we want to use it as a handler.
package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog Dog Dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	// http.HandleFunc("/dog", d)
	// http.HandleFunc("/cat", c)
	// To make it work with http.handle?
	// http.handle wants a handler.
	// handler is anything that implements
	// ServeHTTP(w ResponseWriter, r *Request)
	// HandlerFunc has the ServeHTTP method attached to it -
	// https://pkg.go.dev/net/http#HandlerFunc.ServeHTTP
	// which makes HandlerFunc a handler, which is what Handle wants.
	// HandlerFunc's underlying type is - func(ResponseWriter, *Request)
	// func 'd' and 'c' are the functions with the same parameters, which is
	// exactly the underlying type of HandlerFunc
	// So we can use convert d and c to Handlerfunc. CONVERSION.
	// we are converting d and c to type HandlerFunc.
	// And HandlerFunc has the ServeHTTP method that means HandlerFunc is implementing handler interface
	// which is what http.Handle wants as it's second argument.
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}

// type hotdog int

// func main(){
// 	var x hotdog
// 	x = 7
// 	var y int
//  y = x (This does not work)
// // but if we convert x to an int, since x's underlying type is int.
// // We can convert it and then assign it.
// 	y = int(x)
// 	fmt.Println(7)
// }
