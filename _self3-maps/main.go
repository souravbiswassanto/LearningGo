package main

import (
	"fmt"
)

func main() {
	Map := make(map[string]int)
	Map["sourav"] = 1
	Map["santo"] = 2

	for key, value := range Map {
		fmt.Println(key, value)
	}
	val, ok := Map["sato"]
	fmt.Println(val, ok)
	if ok == true {
		fmt.Println("okay")
	}

	// creating more complex maps

	newmap := make(map[string]map[string]int)
	innermap := make(map[string]int)
	innermap["santo"] = 1
	newmap["sourav"] = innermap

	fmt.Println(newmap)
	for key1, value := range newmap {
		for key2, val := range value {
			fmt.Println(key1, key2, val)
		}
	}
}
