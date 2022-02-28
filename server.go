package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"unicode"
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
	illegalChar := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	fields := strings.FieldsFunc(message, splitString)

	if len(fields) == 2 || len(fields) == 3 {
		if fields[0] == "INDEX" || fields[0] == "REMOVE" || fields[0] == "QUERY" {
			payload := fields[1]
			if strings.IndexFunc(payload, illegalChar) != -1 {
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
		//Read from TCP
		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		//Remove whitespaces and split strings
		temp := strings.TrimSpace(string(message))
		fmt.Println(temp)

		//Check Format
		if checkFormat(temp) {
			splitString := func(c rune) bool {
				return c == '|'
			}
			fields := strings.FieldsFunc(message, splitString)
			command := fields[0]
			if len(fields) == 2 || len(fields) == 3 {
				if command == "INDEX" {
					c.Write([]byte("OK\n"))
				}
				if command == "REMOVE" {
					c.Write([]byte("OK\n"))
				}
				if command == "QUERY" {
					c.Write([]byte("OK\n"))
				}
			}
		} else {
			c.Write([]byte("ERROR\n"))
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
