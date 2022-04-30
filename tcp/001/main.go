package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)

	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		io.WriteString(conn, "\n helloo form tcp")
		fmt.Fprintln(conn, "\nhow is ur day")
		fmt.Fprintf(conn, "\n%v", "i hope ur well")
		conn.Close()

	}

}
