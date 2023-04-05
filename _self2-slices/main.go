package main

import "fmt"

func main() {
	vector := make([]int, 1)
	vector = append(vector, 1)
	vector = append(vector, 2)
	//print(vector)
	for _, value := range vector {
		fmt.Println(value)
	}

}
