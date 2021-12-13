// Notice how when we import we import the package "go.uuid"
// while when we use the package in the code we use uuid.NewV4()
// because for go.uuid the folder name is not the same as the package name.
// That's a convention in go to make the folder name same as the package name.
// satori didn't follow that.
// If we don't want the program to be cookie dependednt we can link it to the every link
// on our web page.
// So for every link people clicked on our site there would be a unique id sent
// to the server as a url parameter.
package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8086", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	// we call the cookie, we pass in the string, name of the cookie.
	cookie, err := req.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4() // There are many types of uuids and NewV4 doesn't take in any parameters.
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(), // id.String() method just converts the uuid into string.
			// Secure: true,
			// If we were doing https we would have the above line. Which means that this cookie
			// is only accessible when the session is secure. So the cookie would only be sent
			// if we were on the https.
			HttpOnly: true, // This means we cannot access this cookie via Javascript which makes
			// this cookie more secure.
		}
		http.SetCookie(res, cookie)
	}
	fmt.Fprintln(res, cookie)
}
