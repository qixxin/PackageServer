package main

import "fmt"

type Package struct {
	name       string
	dependency string
}

var packageList map[string]Package

func main() {
	//message := "INDEX|cloog|gmp,isl,pkg-config\n"
	/* message := "INDEXclooggmp,isl,pkg-config\n"
	temp := strings.TrimSpace(string(message))
	fmt.Println(temp)
	checkFormat := func(c rune) bool {
		return c == '|'
	}
	fields := strings.FieldsFunc(temp, checkFormat)
	fmt.Println(fields)
	fmt.Println(len(fields)) */
	packageList = make(map[string]Package)
	packageList["package1"] = Package{}
	fmt.Println("Hello World")
}
