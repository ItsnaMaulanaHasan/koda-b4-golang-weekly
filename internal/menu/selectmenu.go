package menu

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/cart"
	"os"
	"strconv"
	"strings"
)

func SelectMenu() {
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	for loop {
		fmt.Printf("\x1bc")
		fmt.Print("--- Select Menu ---\n\n")

		for _, item := range Menus {
			fmt.Printf("%d. %s - Rp %.2f\n", item.ID, item.Name, item.Price)
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
		for _, item := range Menus {
			if item.ID == choice {
				fmt.Println("\nYou selected:", item.Name)
				found = true
				itemExists := false
				for i := range cart.Carts {
					if cart.Carts[i].ID == item.ID {
						cart.Carts[i].Quantity++
						itemExists = true
						break
					}
				}

				if !itemExists {
					cart.Carts = append(cart.Carts, cart.CartItem{
						ID:       item.ID,
						Name:     item.Name,
						Quantity: 1,
						Price:    item.Price,
					})
				}
				fmt.Printf("%s has been added to your cart.\n", item.Name)
			}
		}

		if !found {
			fmt.Print("Menu item not found, please try again... ")
			scanner.Scan()
			continue
		}

		fmt.Print("\nPress Enter to continue... ")
		scanner.Scan()
	}
}
