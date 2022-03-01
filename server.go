package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"unicode"
)

const (
	host = "localhost"
	port = "8080"
)

var packageList = struct {
	sync.RWMutex
	m map[string]map[string]string
}{m: make(map[string]map[string]string)}

//var packageList = make(map[string]map[string]string)
var mutex = &sync.RWMutex{}

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

//Check if dependencies are indexed
func dependenciesCheck(message []string) bool {
	mutex.RLock()
	for i := 1; i < len(message)-1; i++ {
		if _, ok := packageList.m[message[i]]; ok {
			mutex.RUnlock()
			return true
		}
	}
	mutex.RUnlock()
	return false
}

func removalDependenciesCheck(message string) bool {
	mutex.RLock()
	for _, value := range packageList.m {
		if len(value) != 0 {
			for _, dependency := range value {
				if dependency == message {
					mutex.RUnlock()
					return true
				}
			}
		}
	}
	mutex.RUnlock()
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
		fmt.Println(checkFormat(temp))
		//Command logic
		if checkFormat(temp) {
			splitString := func(c rune) bool {
				return c == '|'
			}
			splitDependencies := func(c rune) bool {
				return c == ','
			}
			fields := strings.FieldsFunc(message, splitString)
			messageLength := len(fields)
			command := fields[0]
			packageName := fields[1]
			if messageLength == 2 || messageLength == 3 {
				if command == "INDEX" {
					if messageLength == 2 {
						if _, ok := packageList.m[packageName]; ok {
							delete(packageList.m, packageName)
						}
						packageList.m[packageName] = map[string]string{}
						c.Write([]byte("OK\n"))
					} else {
						dependencies := strings.FieldsFunc(fields[2], splitDependencies)
						if dependenciesCheck(dependencies) {
							if _, ok := packageList.m[packageName]; ok {
								delete(packageList.m, packageName)
								packageList.m[packageName] = map[string]string{}
								for i := 1; i < len(dependencies)-1; i++ {
									packageList.m[packageName][dependencies[i]] = dependencies[i]
								}
								c.Write([]byte("OK\n"))
							} else {
								packageList.m[packageName] = map[string]string{}
								for i := 1; i < len(dependencies)-1; i++ {
									packageList.m[packageName][dependencies[i]] = dependencies[i]
								}
								c.Write([]byte("OK\n"))
							}
						} else {
							c.Write([]byte("FAIL\n"))
						}
					}
				}
				if command == "REMOVE" {
					//c.Write([]byte("OK\n"))
					mutex.RLock()
					if _, ok := packageList.m[packageName]; ok {
						mutex.RUnlock()
						if removalDependenciesCheck(packageName) {
							c.Write([]byte("FAIL\n"))
						} else {
							mutex.Lock()
							delete(packageList.m, packageName)
							mutex.Unlock()
							c.Write([]byte("OK\n"))
						}
					} else {
						//c.Write([]byte("OK\n"))
						_, err := c.Write([]byte("OK\n"))
						if err != nil {
							fmt.Println(err)
							return
						}
					}

				}
				if command == "QUERY" {
					if _, ok := packageList.m[command]; ok {
						c.Write([]byte("OK\n"))
					} else {
						c.Write([]byte("FAIL\n"))
					}
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
	l, err := net.Listen("tcp", host+":"+port)
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
