// Now we'll see how to write code, to create a client to dial into our TCP server.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8086")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs))
}
