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
		log.Fatal("error", "error", err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("error", "error", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input:", err)
		}
		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("error sending data:", err)
		}
	}
}
