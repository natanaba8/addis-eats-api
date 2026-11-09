package main

import "fmt"

func main() {

	var name string = "Hello, World!"
	var age int = 30

	address := "Addis Ababa, Ethiopia"

	var (
		_ string = "John Doe"
		_ int    = 25
		_ string = "addis ababa, Ethiopia"
	)

	type addressType string

	const (
		User int = iota
		Editor
		Admin
	)

	fmt.Println("Your name is:", name)
	fmt.Println("your age is:", age)
	fmt.Println("and your address is:", address)
	fmt.Println("Your name is:", name)
	fmt.Println("Your role is:", User)
}
