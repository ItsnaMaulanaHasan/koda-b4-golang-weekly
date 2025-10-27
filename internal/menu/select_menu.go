package menu

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/models"
	"golang-weekly/internal/utils"
	"text/tabwriter"
)

func SelectMenu(reader *bufio.Reader, scanner *bufio.Scanner, w *tabwriter.Writer) {
	loop := true
	for loop {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Print(r)
					scanner.Scan()
				}
			}()
			fmt.Printf("\x1bc")
			fmt.Print("----------------- Select Menu -----------------\n\n")

			fmt.Println("-----------------------------------------------")
			fmt.Fprintln(w, "No\tName\tPrice")
			fmt.Fprintln(w, "---\t----------------------------\t------------")

			models.PrintRows(&models.Menu{}, w)

			w.Flush()

			fmt.Print("-----------------------------------------------\n")

			fmt.Print("\n0. Back to Home\n")
			fmt.Print("\nChoose a menu: ")

			choice, err := utils.InputInt(reader)

			if err != nil {
				panic("Invalid input, please enter a number... ")
			}

			if choice == 0 {
				loop = false
				return
			}

			found := false
			for _, item := range models.Menus {
				if item.ID == choice {
					fmt.Println("\nYou selected:", item.Name)
					found = true
					itemExists := false
					for i := range models.CartOrders.ListCart {
						if models.CartOrders.ListCart[i].ID == item.ID {
							models.CartOrders.ListCart[i].Quantity++
							itemExists = true
							break
						}
					}

					if !itemExists {
						models.CartOrders.ListCart = append(models.CartOrders.ListCart, models.CartItem{
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
				panic("Menu not found, please try again... ")
			}

			fmt.Print("\nPress Enter to continue... ")
			scanner.Scan()
		}()
	}
}
