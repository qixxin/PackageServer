package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"unicode"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	message1 := strings.TrimSpace(string(buf))
	// Send a response back to person contacting us.
	fmt.Println(message1)

	splitString := func(c rune) bool {
		return c == '|'
	}

	fields := strings.FieldsFunc(message1, splitString)
	messageLength := len(fields)
	command := fields[0]
	packages := fields[2]
	fmt.Println(messageLength)
	fmt.Println(command)
	fmt.Println(packages)
	/* if checkFormat(message1) {
		splitString := func(c rune) bool {
			return c == '|'
		}

		fields := strings.FieldsFunc(message1, splitString)
		messageLength := len(fields)
		command := fields[0]
		if messageLength == 2 || messageLength == 3 {
			if command == "INDEX" {

			}
			if command == "REMOVE" {
				conn.Write([]byte("OK\n"))
			}
			if command == "QUERY" {

			}
		}
	} else {
		conn.Write([]byte("ERROR\n"))
	} */
}

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
