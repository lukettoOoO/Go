// the net package provides the tools we need to easily build tcp socket servers

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// net.Listen starts the server on the given network (tcp) and address (port 8090 on all interfaces)
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	// close the listener to free the port when the application exits
	defer listener.Close()

	// loop indefinitely to accept new client connections
	for {
		// wait for a connection
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting conn:", err)
			continue
		}

		// we use a goroutine here to handle the connection so that the main loop can continue accepting more connections
		go handleConnection(conn)
	}
}

// handleConnection handles a single client connection, reading one line of text from the client and returning a response
func handleConnection(conn net.Conn) {
	// closing the connection releases resources when we are finished interacting with the client
	defer conn.Close()

	// use bufio.NewReader to read one line of data from the client (terminated by a newline)
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}

	// create and send a response back to the client, demonstrating two-way communication
	ackMsg := strings.ToUpper(strings.TrimSpace(message))
	response := fmt.Sprintf("ACK: %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server write error: %v", err)
	}
}

// run the tcp server in the background
// send data and capture the response using netcat
