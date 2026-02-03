package main

import (
	"fmt"
	"log"
	"net"

	"github.com/voznyibohdan/httpfromtcp/internal/request"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())

		req, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}

		fmt.Printf(
			"Request line:\n- Method: %s\n- Target: %s\n- Version: %s\n",
			req.RequestLine.Method,
			req.RequestLine.RequestTarget,
			req.RequestLine.HttpVersion,
		)
	}
}
