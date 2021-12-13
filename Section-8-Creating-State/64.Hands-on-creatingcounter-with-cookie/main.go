// Create a program that will use a cookie to track how many times a user has been to the website/domain.

package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	// We need to handle our /favicon.ico otherwise our favicon.ico will go to "/" route because "/"
	// catches everything which will incrememnt our counter everytime /favicon.ico comes.
	// So our "/" catches everything except "/favicon.ico".
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8086", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")

	// ErrNoCookie is returned by Request's Cookie method when a cookie is not found.
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{ // we need a pointer because SetCookie takes in a pointer.
			Name:  "my-cookie",
			Value: "0",
		}
	}
	count, err := strconv.Atoi(cookie.Value) //string ocnvert ASCII to int.
	if err != nil {
		log.Fatalln(err) // we cannot increment the count if the count is not there hence we did a "Fatalln" here.
	}
	count++
	cookie.Value = strconv.Itoa(count) // Convert the count back to ASCII.

	http.SetCookie(res, cookie)

	io.WriteString(res, cookie.Value)

	// Note: This is not the best counter as someone could delte the cookie
	// from the browser and our count will be back to 1.
}
