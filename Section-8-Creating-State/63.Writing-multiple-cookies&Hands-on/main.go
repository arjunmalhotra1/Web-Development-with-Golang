// We can also write multiple cookies to the client's machine.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// We can write multiple cookies to client's machine.

func main() {
	http.HandleFunc("/", set)                // Writes one cookie.
	http.HandleFunc("/read", read)           // Reads all the cookies.
	http.HandleFunc("/abundance", abundance) // Writes a bunch of cookies.
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8086", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "some-value",
	})
	fmt.Fprintln(res, "Cookie written check your browser")
	fmt.Fprintln(res, "in chrome go to dev tools / application / cookies")
}

func read(res http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #1:", c1)
	}

	c2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #2:", c2)
	}

	c3, err := req.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #3:", c3)
	}
}

func abundance(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "general",
		Value: "Some other value about general things",
	})
	http.SetCookie(res, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific thing",
	})
	fmt.Fprintln(res, "Cookies written - check your browser")
	fmt.Fprintln(res, "")
}
