package cart

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/models"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

var CartMenus = []models.MenusPage{
	{ID: 1, Menu: "Checkout", Action: CheckoutCart},
	{ID: 2, Menu: "Edit Cart", Action: EditCart},
	{ID: 3, Menu: "Clear Cart", Action: ClearCart},
}

func showCarts() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	fmt.Println("----------------------------------------------------")
	fmt.Fprintln(w, "No\tName\tQty\tSubtotal")
	fmt.Fprintln(w, "---\t----------------------------\t---\t------------")

	models.PrintRows(models.CartOrders, w)

	w.Flush()

	fmt.Print("----------------------------------------------------\n")
	fmt.Printf("Total\t\t\t\t        Rp.%.2f", getTotal(models.CartOrders.ListCart))
	fmt.Print("\n----------------------------------------------------\n\n")
}

func CartsPage() {
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	for loop {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Print(r)
					scanner.Scan()
				}
			}()
			fmt.Println("\x1bc")
			fmt.Print("----------------- Your Carts -----------------------\n\n")

			if len(models.CartOrders.ListCart) == 0 {
				fmt.Print("Your carts is empty.\n\n")
				fmt.Print("----------------------------------------------------\n\n")
				fmt.Print("Press enter to go back to the main menu... ")
				scanner.Scan()
				loop = false
				return
			}

			showCarts()

			for _, menu := range CartMenus {
				fmt.Printf("%d. %s\n", menu.ID, menu.Menu)
			}

			fmt.Print("\n0. Exit\n")

			fmt.Print("\nChoose a menu: ")

			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)
			choice, err := strconv.Atoi(choiceStr)

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
						menu.Action()
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
