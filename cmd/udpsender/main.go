package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
		}

		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Print(err)
		}
	}

	// to create a udp listener: nc -u -l 42069
}
