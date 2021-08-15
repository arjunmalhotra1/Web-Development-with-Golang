package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8086")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
