package main

import "fmt"

type set map[string]string

func call(ls set) set {
	return ls
}

func caller() set {
	val := mos()
	val["monu"] = "good"
	return call(val)
}

func mos() map[string]string {
	return map[string]string{
		"hello": "world",
		"echo":  "rush",
	}
}

func (r set) a(ls string) set {
	r[ls] = "av"
	return r
}

func b(m map[string]string) map[string]string {
	m["monu"] = "good"
	return m
}

func main() {
	fmt.Println(caller())
	//av := make(set)
	av := set(map[string]string{
		"a": "b",
	}).a("d")
	ad := set(map[string]string{
		"accc": "b",
	}).a("dccc")

	fmt.Println(av.a("str").a("av"))
	fmt.Println(ad.a("strrrrr").a("aadadadv"))
	fmt.Printf("%T\n", ad)
}
