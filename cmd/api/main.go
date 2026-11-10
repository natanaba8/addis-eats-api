package main

import "fmt"

func test(param1 int) (int, string) {
	var err string
	if param1%2 != 0 {
		err = "input is odd"
	}
	return param1, err
}

func main() {
	// _, err := test(100)

	x := 100
	y := x

	fmt.Println(y)

	fmt.Println(y)

	result, err := test(y)
	fmt.Println("Result:", result, "Error:", err)
}

// func main() {

// 	var name string = "Hello, World!"
// 	var age int = 30

// 	address := "Addis Ababa, Ethiopia"

// 	var (
// 		name    string = "John Doe"
// 		age     int    = 25
// 		address string = "addis ababa, Ethiopia"
// 	)

// 	type addressType string

// 	const (
// 		User int = iota
// 		Editor
// 		Admin
// 	)

// 	fmt.Println("Your name is:", name)
// 	fmt.Println("your age is:", age)
// 	fmt.Println("and your address is:", address)
// 	fmt.Println("Your name is:", name)
// 	fmt.Println("Your role is:", User)
// 	fmt.Println("My Statuse:", name)
// 	fmt.Println("My Statuse:", age)
// 	fmt.Println("My Statuse:", address)
// }
