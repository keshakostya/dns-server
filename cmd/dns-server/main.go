package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		fmt.Println("Error creating socket: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("Listening on 8080 port")

	buf := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			fmt.Println("Error reading datagram: ", err)
		}

		fmt.Printf("Received %d bytes from %s\n", n, addr)
		nlCounter := 0

		fmt.Printf("======PKT START======\n")
		for _, b := range buf[:n] {
			nlCounter++
			fmt.Printf("%x ", b)
			if nlCounter == 15 {
				fmt.Printf("\n")
				nlCounter = 0
			}
		}
		if nlCounter < 15 {
			fmt.Printf("\n")
		}
		fmt.Printf("======PKT END======\n")
	}
}
