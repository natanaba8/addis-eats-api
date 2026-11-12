package main

import (
	"fmt"
)

func main() {
	// const SantimPerBirr = 100
	// const ServiceRatePct float64 = 5

	// type Status int

	// const (
	// 	Placed Status = iota
	// 	Accepted
	// 	Cooking
	// 	onTheway
	// 	Delivered
	// )

	// fmt.Println("ServiceRate :", ServiceRatePct)

	// fmt.Println("Statuses:")
	// fmt.Println("  Placed:", Placed)
	// fmt.Println("  Accepted:", Accepted)
	// fmt.Println("  Cooking:", Cooking)
	// fmt.Println("  On The Way:", onTheway)
	// fmt.Println("  Delivered:", Delivered)

	// feeSantim := 500
	// const area = ""

	// if area == "Megenagna" {
	// 	feeSantim := 4500
	// 	fmt.Println(feeSantim)
	// }

	// fmt.Println(feeSantim)

	// if area == "Megenagna" {
	// 	feeSantim = 4500
	// }
	// fmt.Println(feeSantim)

	// _, err := strconv.Atoi(input)
	// for _, dish := range menu { ... }

	// var subtotalSantim int = 68500
	// var ratePct float64 = 5

	// service := int(float64(subtotalSantim) * ratePct / 100)
	// total := subtotalSantim + service

	// fmt.Println(total)

	// var small int32 = 900
	// var big int64 = int64(small)

	// fmt.Println(big)

	dish := "ዶሮ ወጥ"

	fmt.Println(len(dish))
	fmt.Println(len([]rune(dish)))

	for i, r := range dish {
		fmt.Printf("%d %c", i, r)
	}

	// name := "Yod Abyssinia"
	// fmt.Println(string.ToUpper(name))
	// fmt.Println(string.HasPrefix("0911223344", "09"))
}
