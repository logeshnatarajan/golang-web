package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {

		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}

}
func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)
	// It will keep on listening untill the network manually breaks
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "i heared you saying %s\n", ln)
	}
	defer conn.Close()
	fmt.Println("djhfery")
}
