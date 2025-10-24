package cart

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowCarts() {
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	for loop {
		fmt.Println("\x1bc")
		fmt.Print("--- Your Cart ---\n\n")

		if len(Carts) == 0 {
			fmt.Print("Your cart is empty.\n\n")
			fmt.Print("-----------------\n\n")
			fmt.Print("Enter to go back to the main menu... ")
			scanner.Scan()
			loop = false
		} else {
			total := 0.0
			for _, item := range Carts {
				subtotal := float64(item.Quantity) * item.Price
				fmt.Println(item.Name, "- Quantity:", item.Quantity, "- Subtotal: Rp", subtotal)
				total += subtotal
			}

			fmt.Println("\nTotal: Rp", total)
			fmt.Print("\n-----------------\n\n")

			for _, menu := range CartMenus {
				fmt.Printf("%d. %s\n", menu.ID, menu.Menu)
			}

			fmt.Print("\n0. exit\n\n")

			fmt.Print("\nChoose a menu: ")

			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)
			choice, err := strconv.Atoi(choiceStr)

			if err != nil {
				fmt.Print("Invalid input, please enter a number... ")
				scanner.Scan()
				continue
			}

			if choice == 0 {
				loop = false
				break
			}

			found := false
			for _, menu := range CartMenus {
				if menu.ID == choice {
					if menu.Action != nil {
						menu.Action()
					} else {
						fmt.Print("Menu action not implemented yet...")
						scanner.Scan()
					}
					found = true
					break
				}
			}

			if !found {
				fmt.Print("Invalid menu option, press enter to continue...")
				scanner.Scan()
			}
		}
	}
}
