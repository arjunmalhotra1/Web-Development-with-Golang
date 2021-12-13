// Here we will upload a file.
// We will haev a form that will aloow a user to upload a file.
// We will have code in our server to deal with the uploaded file.
// In first example we read the file and wrtite the content of the file back to the client.
// In second example we take the content of the first file and write it into the second file and
// store it on the server.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	// He said we have to handle our "favicon". As we don't want to read our request from "favicon"
	// and handle it as a file. So if a request comes through the "/favicon.ico" route we just say
	// we don't have that, it's not found. It's just redundant code, he wanted to demonstrate the
	// use of http.NotFoundHandler.
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8086", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	// http.MethodPostis a constant and it's better to use constants.
	if req.Method == http.MethodPost {
		//open
		f, h, err := req.FormFile("q") // Here we not FormValue but FormFile.
		// FormFile is used for catching a file submitted by a user.
		// Here formFile catches a file, as it's got the identifier "q"
		// FormFile returns us, "(multipart.File, *multipart.FileHeader, error)"
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// for your information
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// read
		// ReadAll takes an io.Reader and gives us back a slice of byte and an error.
		// We can pass in a file to ReadAll, as ReadAll wants a reader.
		// "f" returned by FormFile is a multipart.File and "File" is an interface with io.Reader method.
		// We finaly convert the slice of bytes into a "string"
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}
	// When the controlles comes to the default route "/"
	// If the method is "POST" we will run the dorm submisison code within line 32 and 53.
	// If the method is either POST or GET either way lines 58 to 64 run.
	// We set the response header to "text/html" and then it writes back to the response the form.
	// Form is for uploading a file and writes "s". "s" is either going to be an empty string
	// since "var s string" initializes "s" to the zero value OR
	// if the method is POST and that logic ran, it will be the contents of the files were uploaded.
	// enctype="multipart/form-data This just says that we are uploading the form data.
	// <input type="file" name="q"> gives us the "Choose File" button.
	// We could have had
	// <form action ="/" method="POST" enctype="multipart/form-data">
	// That would send it back to "/". but we don't need it here as the response goes back to the same page in our case.
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
