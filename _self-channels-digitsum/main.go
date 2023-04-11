package main

import (
	"bufio"
	. "fmt"
	"os"
)

/*
This post explores four of the less common properties of channels:

    A send to a nil channel blocks forever
    A receive from a nil channel blocks forever
    A send to a closed channel panics
    A receive from a closed channel returns the zero value immediately
*/

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func digits(number int, digch chan int) {
	for number != 0 {
		singledigit := number % 10
		digch <- singledigit
		number /= 10
	}
	close(digch)
}

func squaresum(number int, sqrch chan int) {
	sqdgch := make(chan int)
	go digits(number, sqdgch)
	sum := 0
	for v := range sqdgch {
		sum += v * v
	}
	sqrch <- sum
}

func cubesum(number int, cubech chan int) {
	dgch := make(chan int)
	go digits(number, dgch)
	sum := 0
	for v := range dgch {
		sum += v * v * v
	}
	cubech <- sum
}

func main() {
	defer out.Flush()
	var number int
	Fscan(in, &number)
	sqrch := make(chan int)
	cubech := make(chan int)
	go squaresum(number, sqrch)
	go cubesum(number, cubech)
	cubesum, squaresum := <-cubech, <-sqrch
	Fprintf(out, "%v\n", cubesum+squaresum)
}
