package main

import "fmt"

func main() {
	const SantimPerBirr = 100
	const ServiceRatePct float64 = 5

	type Status int

	const (
		Placed Status = iota
		Accepted
		Cooking
		onTheway
		Delivered
	)

	fmt.Println("Statuses:")
	fmt.Println("  Placed:", Placed)
	fmt.Println("  Accepted:", Accepted)
	fmt.Println("  Cooking:", Cooking)
	fmt.Println("  On The Way:", onTheway)
	fmt.Println("  Delivered:", Delivered)
}
