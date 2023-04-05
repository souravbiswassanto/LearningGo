package main

import "fmt"

func f(list []float32) float32 {
	tot := float32(0.0)
	for _, val := range list {

		tot += val
	}
	return tot / float32(len(list))
}

func main() {
	list := make([]float32, 0)
	fmt.Println(list)
	list = append(list, 3.5)
	list = append(list, 5.5)
	list = append(list, 2.55)
	tot := float32(0.0)
	for _, val := range list {
		fmt.Println(val)
		tot += val
	}
	fmt.Println(tot / float32(len(list)))
	fmt.Println(f(list))
}
