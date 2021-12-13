// Redirects allow us to redirect a request from one location to another.
// When a client makes a request tot he server, the client is asking for a particular resource,
// which is done via the Uniform Resource Locator or the Uniform Resource Identifier.
// Server can choose to redirect to another location.
// May be that resource is moved to another location or may be some other processing.
// Client makes a request to a server and the request if for certain resource
// and the server just redirects the client to another location.
// Status codes 300s are the redirects.
// https://pkg.go.dev/net/http#pkg-constants
// We will learn about
// 301 - Status moved permanently.
// 303 - Status see Other &
// 307 - Temporary Redirect.
// HTTP - Defines the rules of communication.
// HTTP specifies how the text should be formatted when we send the data between the client and Server.
// It has the request line and the status line when it's formatting.

// Both the request & response have a start line, headers, and a body.
// The request start line is called the “request line”.
// It consists of
// 			request-line = method SP request-target SP HTTP-version CRLF
// method is GET/POST
// request-target is some URL.Give me this URL or give me that URL.
// Then the http-version.

// The response start line is called the “status line”. It consists of
// 			status-line = HTTP-version SP status-code SP reason-phrase CRLF
// Response/status-line is the HTTP-version
// then a status code - 200s/300s/400s
// and a reason-phrase.

// 301 is moved permanently.
// If you aske the browser to fetch a URL and if it returns 301, the browser remembers it and
// the browser will never hit the server again for that URL.
// Say in the browser we ask for (resource) /dog and the server responds with moved permanently.
// And the server writes the header that puts the location for "/doggy"
// The client/browser will automatically ask for to give the resource at "/doggy"
// The server will give "/doggy" back.
// The browser/client will remember that and if ever the user wants "/dog" again
// the browser doesn't even go to the server. The browser rememebers that it
// was permanently moved. The "/dog" is moved to "/doggy".

// 303 - See other, always changes the method to GET. Pretty much like,
// Hey! get this for me. I just want to see other.

// 307 - Temporary Redirect - Keeps the same method. Will preserve the method.
// Whatever the thing was normally done, whether the method GET/POST we (server) will just temporary redirect
// to some other location.
// So if we POST to a server with a request, and if the server redirects then that redirect request will
// alos be POST.

// If we POSTed to some location we can 303 to a success page(redirect to the success page).
// Redirect to the success page after the processing has occurred.

// There was 302 which was supposed to preserve otherwise original request, but at a different location.
// Say if we are making POST it will make POST to another location &
// If we are making GET it will make GET to another location.
// But people started using 302 to always be GET request.
// 302 is still there and can be used but it's only for legacy applications.
// WE should use 303 and 307. They bring more clarity.
