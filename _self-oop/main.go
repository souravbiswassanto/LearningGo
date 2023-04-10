package main

import . "fmt"

type Rectangle struct {
	Name          string
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

/*
This can be read as
class Rectangle
  field Name: string
  field Width: float64
  field Height: float64
  method Area() //non-virtual
     return this.Width * this.Height

*/

func main() {
	var a Rectangle
	var b = Rectangle{"I'm b.", 10, 20}
	var c = Rectangle{Height: 12, Width: 14}

	Println(a)
	Println(b)
	Println(c)
}
