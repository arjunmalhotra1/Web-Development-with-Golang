// There might be situation where we would want to get rid of the cookie.
// MaxAge - We set the max age, which is some amount in seconds. When
// we set it to 0 or a negative number then age of that cookie is done.
// We can also set the max age to last a certain period of time. So after a certain
// period of time the session expires.
// Say we would want 10 minutes, if the cookies hasn't been reset/activity at which point
// the session is over.
// Besides MaxAge we could use "Expires" but that is deprecated.
// Expires sets the exact time when the cookie expires. Expires is 'Deprecated'.

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)        // No cookie
	http.HandleFunc("/set", set)       // Set a cookie
	http.HandleFunc("/read", read)     // Read a cookie
	http.HandleFunc("/expire", expire) // Expire a cookie
	http.ListenAndServe(":8086", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, `<h1><a href="/set"> Set a Cookie </a></h1>`)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "session",
		Value: "some-value",
	})
	fmt.Fprintln(res, `<h1><a href="/read"> Read </a></h1>`)
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session") //re.Cookie() takes in a string and retuns a pointer to the cookie & an error.
	if err != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return // remember to have the return here. If we don't return from here code will continue to run through.
	}
	fmt.Fprintf(res, `<h1> Your cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return // If we don't return from here, the "/set" will be called and
	}
	c.MaxAge = -1          // delete cookie
	http.SetCookie(res, c) // Since have changed one of the fields of the cookie, we need to set the cookie.
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
