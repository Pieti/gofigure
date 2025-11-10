package main

import (
	"fmt"
	"log"
	"net"

	"github.com/pieti/httpfromscratch/internal/request"
)

func main() {
	l, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}
		r, err := request.RequestFromReader(conn)
		if err != nil {
			log.Println("error", "error", err)
			conn.Close()
			continue
		}
		fmt.Printf("Request line: \n- Method: %s\n- Target: %s\n- Version: %s\n", r.RequestLine.Method, r.RequestLine.RequestTarget, r.RequestLine.HttpVersion)
	}
}
