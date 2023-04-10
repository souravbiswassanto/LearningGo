package main

import "fmt"

type base struct {
	name string
	val  int
}

func (r base) temp() {
	fmt.Printf("hello from temp %d\n", r.val)
}

type derive struct {
	base
	data float32
}

func (r derive) temp() {
	fmt.Printf("hello from temp %f\n", r.data)
}

func main() {
	var ret derive = derive{base{"sourav", 4}, 3.4}
	ret.temp()
	ret.base.temp()
	fmt.Printf("")
}
