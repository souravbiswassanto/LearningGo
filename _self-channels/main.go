package main

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		fmt.Printf("calcsquares, %d\n", number)
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		fmt.Printf("calccubes, %d\n", number)
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcCubes(number, cubech)
	go calcSquares(number, sqrch)

	squares := <-sqrch
	fmt.Printf("cal sqr\n")
	cubes := <-cubech
	fmt.Println("Final output", squares+cubes)
}

/*
// This code demonstrates how blocking occour also closing

package main

import (
	"fmt"
)

func producer(chnl chan int) {
	fmt.Printf("up\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("mid1\n")
		//time.Sleep(2 * time.Millisecond)
		chnl <- i
		fmt.Printf("mid2\n")
	}
	fmt.Printf("close\n")
	close(chnl)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}

		fmt.Println("Received ", v, ok)
	}
}
*/

/*
// same code using for loop on channel
package main

import (
    "fmt"
)

func producer(chnl chan int) {
    for i := 0; i < 10; i++ {
        chnl <- i
    }
    close(chnl)
}
func main() {
    ch := make(chan int)
    go producer(ch)
    for v := range ch {
        fmt.Println("Received ",v)
    }
}
*/
