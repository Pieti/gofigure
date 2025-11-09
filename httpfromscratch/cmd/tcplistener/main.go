package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		var line string
		for {
			buf := make([]byte, 8)
			n, err := f.Read(buf)
			if err != nil {
				return
			}

			buf = buf[:n]
			if i := bytes.IndexByte(buf, '\n'); i != -1 {
				line += string(buf[:i])
				buf = buf[i+1:]
				out <- line
				line = ""
			}
			line += string(buf)

			if len(line) != 0 {
				out <- line
			}
		}
	}()
	return out
}

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
		fmt.Println("Accepted connection from", conn.RemoteAddr())
		for line := range getLinesChannel(conn) {
			fmt.Printf("read: %s", line)
		}
	}
}
