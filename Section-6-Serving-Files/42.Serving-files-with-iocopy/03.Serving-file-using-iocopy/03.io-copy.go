package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8086", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html; charset=utf-8")

	// When the toby.jpg is asked for it will run the end point on line 11 and which will trigger
	// the function dogPic

	//source is toby.jpg
	// We could have also have had it this way, with absolute reference.
	// At the root
	// io.WriteString(res, `
	// <img src="/toby.jpg">
	// `)

	io.WriteString(res, `
	<img src="toby.jpg">
	`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	// We open the file.
	// os.Open gives us a pointer to the file and an error.
	// pointer to file, *File
	// has a Read method and a Write method attached to it.
	// hence it implements the reader and the writer interface.
	// because of which we can read and write from a file.
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	defer f.Close()
	// We use io copy to copy from the file to our response writer.
	io.Copy(res, f)
}
