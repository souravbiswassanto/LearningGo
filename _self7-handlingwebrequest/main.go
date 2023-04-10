package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var a []int
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)
	a = append(a, 5)
	a = append(a, 6)
	var b []int = a[2:4]
	var d = []int{7, 8}
	a = d
	var c []int = b[:4]

	fmt.Println(b, c)
	var str []byte
	str = append(str, 65)
	fmt.Println(str)
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	_, err = io.Copy(os.Stdout, file)
	file.Close()
}
