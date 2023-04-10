package main

import (
	. "fmt"
)

type interfacedemo interface {
	method1()
}

// class NamedObj
type NamedObj struct {
	Name string
}

// method show
func (n NamedObj) show() {
	Println(n.Name) // "n" is "this"
}

// class Rectangle
type Rectangle struct {
	NamedObj //inheritance
	//interfacedemo
	Width, Height float64
}

// override method show
func (r Rectangle) show() {
	Println("Rectangle ", r.Name) // "r" is "this"
}

/*
   class NamedObj
      Name: string

      method show
        print this.Name

   class Rectangle
      inherits NamedObj
      field Width: float64
      field Height: float64

      method show //override
        print "Rectangle", this.Name
*/

func main() {
	var a = NamedObj{"Joe"}
	var b = Rectangle{NamedObj{"Richard"}, 10, 20}

	a.show()
	b.show()
}
