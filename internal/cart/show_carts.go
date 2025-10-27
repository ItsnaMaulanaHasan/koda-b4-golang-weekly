package cart

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/models"
	"golang-weekly/internal/utils"
	"text/tabwriter"
)

var CartMenus = []models.MenusPage{
	{ID: 1, Menu: "Checkout", Action: CheckoutCart},
	{ID: 2, Menu: "Edit Cart", Action: EditCart},
	{ID: 3, Menu: "Clear Cart", Action: ClearCart},
}

func printCarts(scanner *bufio.Scanner, w *tabwriter.Writer) bool {
	fmt.Print("----------------- Your Carts -----------------------\n\n")
	if len(models.Carts) == 0 {
		fmt.Print("Your carts is empty.\n\n")
		fmt.Print("----------------------------------------------------\n\n")
		fmt.Print("Press enter to go back to the main menu... ")
		scanner.Scan()
		return false

	}
	fmt.Println("----------------------------------------------------")
	fmt.Fprintln(w, "No\tName\tQty\tSubtotal")
	fmt.Fprintln(w, "---\t----------------------------\t---\t------------")

	models.PrintRows(&models.Cart{}, w)

	w.Flush()

	fmt.Print("----------------------------------------------------\n")
	fmt.Printf("Total\t\t\t\t        Rp.%.2f", getTotal(&models.Carts))
	fmt.Print("\n----------------------------------------------------\n\n")
	return true
}

func ShowCarts(reader *bufio.Reader, scanner *bufio.Scanner, w *tabwriter.Writer) {
	loop := true
	for loop {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Print(r)
					scanner.Scan()
				}
			}()
			fmt.Println("\x1bc")

			loop = printCarts(scanner, w)
			if !loop {
				return
			}

			for _, menu := range CartMenus {
				fmt.Printf("%d. %s\n", menu.ID, menu.Menu)
			}

			fmt.Print("\n0. Exit\n")

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
			for _, menu := range CartMenus {
				if menu.ID == choice {
					if menu.Action != nil {
						menu.Action(reader, scanner, w)
					} else {
						panic("Menu action not implemented yet...")
					}
					found = true
					break
				}
			}

			if !found {
				panic("Invalid menu option, press enter to continue...")
			}
		}()
	}
}
