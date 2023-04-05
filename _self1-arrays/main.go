package main

import "fmt"

func main() {

	var arr [5]int
	arr[4] = 100
	fmt.Print(arr)
	var total int;
	for i, value := range arr {
		total += value
		fmt.Println(arr[i])
	}
	fmt.Println(total)
}