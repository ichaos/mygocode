// Solution to part 3 of the Whispering Gophers code lab.
//
// This program listens on the host and port specified by the -listen flag.
// For each incoming connection, it launches a goroutine that reads and decodes
// JSON-encoded messages from the connection and prints them to standard
// output.
//
// You can test this program by running it in one terminal:
// 	$ part3 -listen=localhost:8000
// And running part2 in another terminal:
// 	$ part2 -dial=localhost:8000
// Lines typed in the second terminal should appear as JSON objects in the
// first terminal.
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
)

var listenAddr = flag.String("listen", "localhost:8000", "host:port to listen on")

type Message struct {
	Body string
}

func main() {
	flag.Parse()
	log.Println("Listen on", *listenAddr)

	// TODO: Create a net.Listener listening from the address in the "listen" flag.
	l, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// TODO: Accept a new connection from the listener.
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go serve(c)
	}
}

func serve(c net.Conn) {
	// TODO: Use defer to Close the connection when this function returns.
	defer c.Close()
	fmt.Fprintln(c, "Welcome!")
	// TODO: Create a new json.Decoder reading from the connection.
	d := json.NewDecoder(c)
	for {
		// TODO: Create an empty message.
		var m Message
		// TODO: Decode a new message into the variable you just created.
		if err := d.Decode(&m); err != nil {
			log.Fatal(err)
		}
		// TODO: Print the message to the standard output.
		fmt.Println("Get from client:", m.Body)
	}
}
