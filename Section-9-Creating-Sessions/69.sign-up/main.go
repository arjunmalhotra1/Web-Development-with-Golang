// To get osmeone to sign up thye have to fill out a form.
// User enters the details and then we process that form.

package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      //userID, user
var dbSessions = map[string]string{} //session ID, UserID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8086", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	// If user is not already loggedin
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	// First thing we check if the user is already logged in.
	// We can distribute the code in the package across multiple files as long as those files
	// are in the same folder.
	// User has to be loggend in to get to the signup page.
	if alreadyLoggedIn(req) { // So if the user is already logged in, you don't need to sign up.
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	//Process Form submission
	if req.Method == http.MethodPost {
		//get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		//username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(res, "Username already taken", http.StatusForbidden)
			return
		}
		//create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = un

		//store user in dbUsers
		u := user{un, p, f, l}
		dbUsers[un] = u

		//redirect
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", nil)
}

// func getUser(res http.ResponseWriter, req *http.Request) user {
// 	// get cookie
// 	cookie, err := req.Cookie("session")
// 	if err != nil {
// 		sID, _ := uuid.NewV4()
// 		cookie = &http.Cookie{
// 			Name:  "session",
// 			Value: sID.String(),
// 		}
// 	}
// 	http.SetCookie(res, cookie)
// 	// if user already exists, get user
// 	var u user
// 	if un, ok := dbSessions[cookie.Value]; ok {
// 		u = dbUsers[un]
// 	}
// 	return u

// }

// func alreadyLoggedIn(req *http.Request) bool {
// 	c, err := req.Cookie("session")
// 	if err != nil {
// 		return false // Return false as user is not currently logged in. User doesn't haev the cookie.
// 	}
// 	un := dbSessions[c.Value] // Grab the username
// 	_, ok := dbUsers[un]      // If dbUsers username is ok, we will return ok.
// 	// "ok" will tell us if the map found the value or if it's jsut returning a zero value.
// 	return ok // If it finds a value it will be true. Else false.
// }
