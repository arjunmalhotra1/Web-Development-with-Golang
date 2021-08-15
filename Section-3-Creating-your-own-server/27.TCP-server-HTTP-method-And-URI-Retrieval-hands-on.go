package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)

	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)

	respond(conn)
}

func request(conn net.Conn) {
	i := 0

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0] // Retrieving the method, GET, POST , PUT, etc.
			u := strings.Fields(ln)[1] // Retrieveing the URI
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
		}
		if ln == "" {
			// Headers are done
			break
		}
		i++

	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html> <html lang ="en"><head><meta
	charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	// This below format is in accordance with the HTTP protocol, of how we should write our response.
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
