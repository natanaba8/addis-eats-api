package main

import "fmt"

// func test(param1 int) (int, string) {
// 	var err string
// 	if param1%2 != 0 {
// 		err = "input is odd"
// 	}
// 	return param1, err
// }

func main() {
	// _, err := test(100)

	// x := 100
	// y := x

	// fmt.Println(y)

	// fmt.Println(y)

	// result, err := test(y)
	// fmt.Println("Result:", result, "Error:", err)

	var restaurant string = "Yod Abyssinia"
	var area = "Bole"
	var priceSantim int

	courier := "dawit tesfaye"
	order, revenue := 12, 84500

	var (
		city     = "Addis Ababa"
		currency = "ETB"
	)

	// var order struct {
	// 	TotalSantim int
	// 	Paid        bool
	// 	Note        string
	// }

	fmt.Println("Restaurant:", restaurant)
	fmt.Println("Area:", area)
	fmt.Println("Price in Santim:", priceSantim)
	fmt.Println("Courier:", courier)
	fmt.Println("Order:", order)
	fmt.Println("Revenue:", revenue)
	fmt.Println("City:", city)
	fmt.Println("Currency:", currency)
	// fmt.Println("*************")
	// fmt.Println("Total Santim: ", order)

	// menu := []string{"ዶሮ ወጥ", "ምሳ አሰራር", "አትክልት አሰራር", "አልጫ ዶሮ", "አይነት", "የተቀለለ እንቁላል"}
	// remaining := 10000
	// servingSantim := 2500
	// done := false

	// for i := 0; i < len(menu); i++ {
	// 	fmt.Println(menu[i])
	// }

	// for remaining > 0 {
	// 	remaining -= servingSantim
	// }

	// for {
	// 	if done {
	// 		break
	// 	}
	// }

	prices := []int{18500, 24000, 12000}

	for i, p := range prices {
		fmt.Println("%d: %.2f ETB\n", i, float64(p)/100)
	}

	fees := map[string]int{
		"Yod Abyssinia": 4500,
		"Addis Ababa":   3500,
		"Bole":          2500,
	}
	for area, fee := range fees {
		fmt.Println(area, fee)
	}

	// for _, p := range prices{
	// 	subtotal += p
	// }

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
