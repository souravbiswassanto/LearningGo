package main

import "fmt"

type User struct {
	data int
}

func (u *User) notify() {
	fmt.Println("hello")
}

func (u *Us) notify() {
	fmt.Println("hello")
}

type demo interface {
	notify()
}

// if v here is a pointer, then receiver can be both pointer and value.

func pnt(v demo) {
	fmt.Printf("%p\n", v)
	v.notify()
}

type Us struct {
}

func main() {
	d := &User{3}
	e := &Us{}
	fmt.Printf("%p\n", d)
	pnt(d)
	pnt(e)
}
