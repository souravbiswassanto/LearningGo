package main

import "fmt"

func main() {
	Map := make(map[string]int)
	Map["sourav"] = 1
	Map["santo"] = 2

	for key, value := range Map {
		fmt.Println(key, value)
	}

}
