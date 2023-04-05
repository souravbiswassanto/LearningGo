package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// this is kind of try catch, if everything goes well then input will hold it
	// else err will catch the errors.
	fmt.Println(input)
	var floatNum float32
	var intNum int32
	fmt.Scanf("%f", &floatNum)
	fmt.Scanf("%d", &intNum)
	fmt.Println(floatNum)
	fmt.Println(intNum)

	i := 1
	for i <= 10 {
		i = i + 1
		fmt.Println(i)
	}

	for i := 1; i <= 10; i++ {
		if i % 2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("even")
		}
	}
}
