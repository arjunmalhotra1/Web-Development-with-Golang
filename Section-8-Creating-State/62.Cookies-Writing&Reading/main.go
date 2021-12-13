// We will write a cookie ot the client's machine
// Cleint makes a request back to our domain, COOKIES ARE DOMAIN SPECIFIC.
// the browser/cleint will see, if there is a cookie specific to our domain.
// Then the client wil send the cookie back to our domain, with the request.
// func SetCookie(res ResponseWriter, cookie *Cookie)
// type Cookie
// func(c *Cookie)String() string // Which allows us to print all the values in the cookie.
// Type Cookie is a struct with many fields and a string method attached to it.
// field "Name" and a field "Value". So we can store whatever we want as strings in the value.
// We could have a UUID or we might do some stuff where we store some other information in value too.

// To get a cookie
// On every request we could pass int he string and get the pointer to the cookie. If there is a cookie present.
// func (r *Request) Cookie(name string)(*Cookie, error)

// To set a cookie, we give a pointer to the cookie.
// func SetCookie(res ResponseWriter, cookie *Cookie)
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8086", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "Some value",
	})
	fmt.Fprintln(res, "Cookie written check your browser")
	fmt.Fprintln(res, "in chrome go to dev tools / application / cookies")
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(res, "Your COOKIE:", c) // since cookie has a string method,
	// this prints out the values of the cookie
}
