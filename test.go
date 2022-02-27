package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "INDEX|cloog|gmp,isl,pkg-config\n"
	temp := strings.TrimSpace(string(message))
	fmt.Println(temp)
	checkFormat := func(c rune) bool {
		return c == '|'
	}
	fields := strings.FieldsFunc(temp, checkFormat)
	fmt.Println(fields)
	fmt.Println(len(fields))
}
