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

func checkFormat(message string) bool {
	splitString := func(c rune) bool {
		return c == '|'
	}
	splitDependencies := func(c rune) bool {
		return c == ','
	}
	fields := strings.FieldsFunc(message, splitString)
	fmt.Println(fields)

	if len(fields) == 2 || len(fields) == 3 {
		if fields[0] == "INDEX" || fields[0] == "REMOVE" {
			payload := fields[1]
			if !strings.Contains(payload, " ") {
				if len(fields) == 2 {
					return true
				} else if len(fields) == 3 {
					dependencies := strings.FieldsFunc(fields[2], splitDependencies)
					for i := 1; i < len(dependencies)-1; i++ {
						if strings.Contains(dependencies[i], " ") {
							return false
						}
					}
					return true
				}
			}
		}
	}
	return false
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		//TCP Read
		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		//Remove whitespaces and split strings
		temp := strings.TrimSpace(string(message))
		fmt.Println(temp)

		if checkFormat(temp) {

		}

		//Initialize TCP writer
		//writer := bufio.NewWriter(c)

		//c.Write([]byte(string(temp)))

	}
	c.Close() // No need, the client Timeout automatically.
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

//First time GOlang ever. First time for everything.
