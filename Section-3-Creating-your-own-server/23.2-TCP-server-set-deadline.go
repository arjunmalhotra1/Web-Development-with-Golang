package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8086")
	if err != nil {
		fmt.Println(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s \n", ln)
	}
	defer conn.Close()

	// Now we get here
	// The connection will timeout
	// that breaks ous our of the scanner loop
	fmt.Println("*** CODE GOT HERE ***")
}
