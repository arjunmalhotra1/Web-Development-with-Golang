https://github.com/GoesToEleven/golang-web-dev/tree/master/017_understanding-net-http-package

type handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// Handler will handle whatever comes in at the address. Then we will be able to make a mux
// which would then do the routing. The mux wil be able to handle different routes.
func ListenAndServe(addr string, handler Handler) error

// For https, secured http
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
