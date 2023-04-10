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

// embedding again

type Discount struct {
	percent   float32
	startTime uint64
	endTime   uint64
}

func (d Discount) Calculate(originalPrice float32) float32 {
	return originalPrice * d.percent
}

func (d *Discount) IsValid(currentTime uint64) bool {
	return currentTime > d.startTime && currentTime < d.endTime
}

type PremiumDiscount struct {
	Discount   //Embedded
	additional float32
}

func (d *PremiumDiscount) Calculate(originalPrice float32) float32 {
	return d.Discount.Calculate(originalPrice) - d.additional
}

func main() {
	var x derived

	fmt.Printf("%T\n", x.a) //=> x.a, float32 (derived.a shadows base.a)

	fmt.Printf("%T\n", x.base.a) //=> x.base.a, string (accessing shadowed member)

	p := Person{Name: "John", Age: 30, Address: Address{Street: "123 Main St", City: "New York", Zip: "10001"}}
	fmt.Println(p.Name, p.Age, p.Street, p.City, p.Zip)

}
