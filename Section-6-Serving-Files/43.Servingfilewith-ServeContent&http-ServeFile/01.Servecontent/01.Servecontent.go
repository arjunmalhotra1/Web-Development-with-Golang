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
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	// When the toby.jpg is asked for it will run the end point on line 11 and which will trigger
	// the function dogPic
	io.WriteString(res, `
	<img src="/toby.jpg">
	`)
}
func dogPic(res http.ResponseWriter, req *http.Request) {
	// f, err := os.Open("/toby.jpg") - I had '/' and it didn't work
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(res, "File not found", 404)
		return
	}

	defer f.Close()
	fi, err := f.Stat() // Stat gives us back FileInfo which is an interface

	// type FileInfo interface {
	// 	Name() string       // base name of the file
	// 	Size() int64        // length in bytes for regular files; system-dependent for others
	// 	Mode() FileMode     // file mode bits
	// 	ModTime() time.Time // modification time
	// 	IsDir() bool        // abbreviation for Mode().IsDir()
	// 	Sys() interface{}   // underlying data source (can return nil)
	// }
	// It has all these above methods and one of which is ModTime() which is
	// Give me the modification time (last itme he file was modified).

	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	//ServeContent asks for
	// ResponseWriter, pointer to a request, file name, modification time and content/actual file
	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)
	// Todd doesn't use ServeContent, with ServeContent we can only serve a single file.
}
