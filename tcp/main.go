package main 

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	ln,err := net.Listen("tcp",":8080")
	if err !=nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
	 	// Reader and writer interfaces very important
		conn,err := ln.Accept()
		if err !=nil {
			//handle error 
		}

	//	Goroutine to handle multiple connections
	go handleConnection(conn)
	}
}



func handleConnection (conn net.Conn) {
	
