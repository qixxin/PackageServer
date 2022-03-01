//To be removed
package main

import (
	"fmt"
	"sync"
)

type Package struct {
	name       string
	dependency string
}

func main() {
	var x = map[string]map[string]string{}
	var mutex = &sync.RWMutex{}
	mutex.Lock()
	x["fruits"] = map[string]string{}
	x["colors"] = map[string]string{}
	//x["cmake"] = map[string]string{}

	x["fruits"]["a"] = "apple"
	x["fruits"]["b"] = "banana"

	x["colors"]["r"] = "red"
	x["colors"]["b"] = "blue"
	mutex.Unlock()
	//x["cmake"][""] = ""
	mutex.RLock()
	if val, ok := x["cmake"]; ok {
		fmt.Println(val)
		if len(val) == 0 {
			fmt.Println("No dependencies")
		}
		fmt.Println("found something")
	}
	mutex.RUnlock()

	//v := x["cmake"]

	//fmt.Println(v)
}
