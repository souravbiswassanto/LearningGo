package main

import (
	"fmt"
)

type base struct {
	a string
	b int
}

type derived struct {
	base // embedding
	d    int
	a    float32 //-SHADOWS
}
type Address struct {
	Street string
	City   string
	Zip    string
}

type Person struct {
	Name    string
	Age     int
	Address // embedding
}

func main() {
	var x derived

	fmt.Printf("%T\n", x.a) //=> x.a, float32 (derived.a shadows base.a)

	fmt.Printf("%T\n", x.base.a) //=> x.base.a, string (accessing shadowed member)

	p := Person{Name: "John", Age: 30, Address: Address{Street: "123 Main St", City: "New York", Zip: "10001"}}
	fmt.Println(p.Name, p.Age, p.Street, p.City, p.Zip)

}
