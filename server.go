package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const (
	host = "localhost"
	port = "8080"
)

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(message))
		c.Write([]byte(string(temp)))
		c.Close()
	}

}

func main() {
	l, err := net.Listen("tcp4", host+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Accept connection on port 8080...")
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}

}
