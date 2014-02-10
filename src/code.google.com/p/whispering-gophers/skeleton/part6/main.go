// Solution to part 6 of the Whispering Gophers code lab.
//
// This program is functionally equivalent to part 5,
// but the reading from standard input and writing to the
// network connection are done by separate goroutines.
//
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"code.google.com/p/whispering-gophers/util"
)

var (
	peerAddr = flag.String("peer", "", "peer host:port")
	self     string
	ch chan Message
)

type Message struct {
	Addr string
	Body string
}

func main() {
	flag.Parse()

	l, err := util.Listen()
	if err != nil {
		log.Fatal(err)
	}
	self = l.Addr().String()
	log.Println("Listening on", self)

	ch = make(chan Message)

	go dial(*peerAddr, ch)
	go readInput(ch)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go serve(c)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	d := json.NewDecoder(c)
	for {
		var m Message
		err := d.Decode(&m)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("%#v\n", m)
	}
}

// TODO: Make a new channel of Messages.
// ch := make(chan Message)

func readInput(ch chan Message) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		m := Message{
			Addr: self,
			Body: s.Text(),
		}
		// TODO: Send the message to the channel of messages.
		ch <- m
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func dial(addr string, ch chan Message) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(addr, err)
		return
	}
	defer c.Close()

	e := json.NewEncoder(c)

	for /* TODO: Receive messages from the channel using range, storing them in the variable m. */ {
		var m = <- ch
		err := e.Encode(m)
		if err != nil {
			log.Println(addr, err)
			return
		}
	}
}
