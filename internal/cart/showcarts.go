package cart

import (
	"fmt"
)

func ShowCarts() {
	loop := true
	for loop {
		fmt.Println("\x1bc")
		fmt.Print("--- Your Cart ---\n")

		if len(Carts) == 0 {
			fmt.Println("Your cart is empty.")
		} else {
			total := 0.0
			for _, item := range Carts {
				subtotal := float64(item.Quantity) * item.Price
				fmt.Println(item.Name, "- Quantity:", item.Quantity, "- Subtotal: Rp", subtotal)
				total += subtotal
			}
			fmt.Println("\nTotal: Rp", total)
		}

		for _, menu := range CartMenus {
			fmt.Printf("%d. %s\n", menu.ID, menu.Menu)
		}

		fmt.Print("\nPress 0 to go back to the main menu ")
		var input string
		fmt.Scanln(&input)
		if input == "0" {
			loop = false
		}
	}
}
