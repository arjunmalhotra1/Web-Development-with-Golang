// To create an HTTP server that responds to an HTTP request.
// We first create a TCP server. This TCP server will handle the request that come in which are
// formatted in a certain way. As when a request comes in that request is just text.
// So if text is formatted in a certain way, if it is adhering to the HTTP (protocols are rules of communication)
// So if the text that comes in to the TCP server adhere s to the HTTP,
// Then our servers will be able to process that text. The standards of a protocol are written in RFC-7230
// RFC - Request for comment. It is a document put out by the IETF - Internet engineering Task Force.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8086") // In the net package there is a function 'Listen'.
	// Listen asks for two strings. First is 'what kind of network do you want to listen on, TCP.
	// Second is what port do you want to listen on.
	// Listen function gives back to us a listener and an err.
	if err != nil {
		log.Panic(err)
	}
	defer li.Close() // We defer the close of the listener, "li"

	for {
		conn, err := li.Accept() // When somebody calls in we accept. This gives us a connection.
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\n Hello from TCP server \n") //WriteString asks for a writer and a string.
		// conn since it implements a write interface is of type writer. We can pass
		// 'conn' to the WriteString method.
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}
