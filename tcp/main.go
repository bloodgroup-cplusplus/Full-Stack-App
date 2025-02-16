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

	defer conn.Close()

	// create a new reader from the connection 

	
	reader := bufio.NewReader(conn)

	// read the command line from the client

	line, err := reader.ReadString('\n')
	if err !=nil {
		fmt.Fprintf(conn,"Error reading command:%v\n",err)
	return 
	}

	// Trim the newline character and split the line into command and resource 
	parts := strings.SplitN(strings.TrimSpace(line), " ",2)
	if len(parts) !=2 {
		fmt.Fprintf(conn, "Invalid command format. Expected format: COMMAND:RESOURCE\n")

	}
}


	
	
