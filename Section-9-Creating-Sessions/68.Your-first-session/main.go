// client send over to the server a unique id.
// Often this unique id is called session id. Usually known as
// "sid". Using the session id there is a table where the
// "session id" is associated with a particular user id.
// So for any given sesison we can see which user is having that session.
// Using hte user id we can pull all the information about a user. As
// every user has a unique user id associated with them.
// A user can have a unique session everytime they log in.
package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{} //User ID, user
// for now we will just have a map for storing the sessions.
// We are using a composite literal here. "{}". But we currently have no values in the dbSessions map.
// Alternative way we could have done that is,
// make(map[string]string)
// we will store the sessionid as the key of the session. The value
// will be the userid
// In the dbUsers "userid" will be the key, and that will give us a user struct.
var dbSessions = map[string]string{} //session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8086", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	// get a cookie
	cookie, err := req.Cookie("session")
	id, _ := uuid.NewV4()
	if err != nil {
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
	}
	http.SetCookie(res, cookie)

	// if the user exista already, get user
	var u user
	// cookie.Value is the session id or unique id which will get us the userid from the dbSessions
	// table.
	// This is ths comma-ok idiom. From a map if you ask for a value. If we give the key and the
	// value is not present then it will just give back the zero value, adn the "ok"
	// variable will be returned as false.
	// "ok {" means if true then we will pull out the user from the dbUsers map.
	// Ans we assign it to "u".
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}
	// Process form submission.
	// If the form is being submitted we can get all these values.
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[cookie.Value] = un // In our case userid is the username.
		dbUsers[un] = u
	}
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

// Inside bar we are just displaying all the user fields attached to
// a user.
func bar(res http.ResponseWriter, req *http.Request) {
	// get cookie.
	cookie, err := req.Cookie("session")
	// In case of an error we redirect it to go set a cookie in "/"
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return // Don't forget to return to stop the code from running
		// all the way down to the rest of the function.
	}
	// This is comma ok idiom again.
	un, ok := dbSessions[cookie.Value] // If no usename with cookie value
	// ok is "false", which means we didn't find a value stored in the dbSessions.
	// Hence we rredirect to "/"
	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(res, "bar.gohtml", u)

}

// map examples with the comma, ok idiom
// https://play.golang.org/p/OKGL6phY_x
// https://play.golang.org/p/yORyGUZviV
