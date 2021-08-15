// To create an HTTP server that responds to an HTTP request.
// We first create a TCP server. This TCP server will handle the request that come in which are
// formatted in a certain way. As when a request comes in that request is just text.
// So if text is formatted in a certain way, if it is adhering to the HTTP (protocols are rules of communication)
// So if the text that comes in to the TCP server adhere s to the HTTP,
// then our servers will be able to process that text. The standards of a protocl are writtn in RFC-7230
// RFC - Request for comment. It is a document put out by the IETF - Internet engineering Task Force.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Simple 3 step process
// 1. We listen.
// 2. We accept a connection, when we accept we get a connection.
// 3. We read or write into that connection.

func main() {
	li, err := net.Listen("tcp", ":8086")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn) // After accepting a connection we launch a goroutine.

	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn) // NewScanner returns a pointer to a Scanner
	for scanner.Scan() {              //scanner.Scan(), Scan() method returna bool.
		// When there is nothing to scan we get a false. That is at the end of the input Scan() returns false.
		// Scan advances the Scanner to the next token, which wil then be available through the Bytes or Text method.
		// Scan scans till a new line '\n' is reached by default.
		// scanner.Split(bufio.ScanWords) this will split on words.
		// scanner.Split(bufio.ScanRunes) A 'Rune' is a character/alphabet in Go.
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// We never get here.
	// We have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")

}
