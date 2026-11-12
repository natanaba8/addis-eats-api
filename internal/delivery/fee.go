package main

import (
	"fmt"
	"strconv"
)

func main() {
	if subtotalSantim >= 50000 {
		feeSantim := 0
	} else if area == "Bole" {
		feeSantim = 3000
	} else {
		feeSantim = 4500
	}

	if qty, err := strconv.Atoi(input); err == nil && qty > 0 {
		order.Quantity = qty
	}

	for i := 0; i < len(menu); i++ {
		fmt.Println(menu[i])
	}

	for remaining > 0 {
		remaining -= servingSantim
	}

	for {
		if done {
			break
		}
	}
}
