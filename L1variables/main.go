package main

import "fmt"

const LoginToken string = "zibirish"

func main() {
	var username string = "saurov"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isOn bool = true
	fmt.Println(isOn)
	fmt.Printf("variable is of type : %T \n", isOn)

	var user float64
	fmt.Println(user)
	fmt.Printf("variable is of type : %T \n", user)
	numberOfUser := 300
	fmt.Println(numberOfUser)
	// defining multiple variables
	var (
		a = 4
		b = 5
		c = 6
	)
	fmt.Println(a, b, c)
}
