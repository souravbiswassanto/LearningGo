package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float32
}

func (c *Circle) area() float32 {
	return c.r * c.r * math.Pi
}

func main() {

	c := Circle{2.0}
	d := Circle{2.5}
	fmt.Println(c.area(), d.area())
	fmt.Println(c.r, d.r)
}
