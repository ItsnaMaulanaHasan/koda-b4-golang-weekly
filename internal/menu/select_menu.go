package menu

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/models"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func SelectMenu() {
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	for loop {
		fmt.Printf("\x1bc")
		fmt.Print("----------------- Select Menu -----------------\n\n")

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

		fmt.Println("-----------------------------------------------")
		fmt.Fprintln(w, "No\tName\tPrice")
		fmt.Fprintln(w, "---\t----------------------------\t------------")

		for _, item := range models.Menus {
			fmt.Fprintf(w, "%d\t%s\tRp.%.2f\n", item.ID, item.Name, item.Price)
		}

		w.Flush()

		fmt.Print("-----------------------------------------------\n")

		fmt.Print("\n0. Back to Home\n")
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
		for _, item := range models.Menus {
			if item.ID == choice {
				fmt.Println("\nYou selected:", item.Name)
				found = true
				itemExists := false
				for i := range models.Carts {
					if models.Carts[i].ID == item.ID {
						models.Carts[i].Quantity++
						itemExists = true
						break
					}
				}

				if !itemExists {
					models.Carts = append(models.Carts, models.CartItem{
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
			fmt.Print("Menu not found, please try again... ")
			scanner.Scan()
			continue
		}

		fmt.Print("\nPress Enter to continue... ")
		scanner.Scan()
	}
}
