package main

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/cart"
	"golang-weekly/internal/history"
	"golang-weekly/internal/menu"
	"golang-weekly/internal/models"
	"golang-weekly/internal/utils"
	"os"
	"text/tabwriter"
)

var HomeMenus = []models.MenusPage{
	{ID: 1, Menu: "Select Menu", Action: menu.SelectMenu},
	{ID: 2, Menu: "Cart", Action: cart.ShowCarts},
	{ID: 3, Menu: "History", Action: history.ShowHistories},
	{ID: 4, Menu: "Clear Cache", Action: utils.ClearCache},
}

func main() {
	defer func() {
		fmt.Printf("\x1bc")
		fmt.Println("\nExiting program")
		os.Exit(0)
	}()
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	for loop {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Print(r)
					scanner.Scan()
				}
			}()
			fmt.Printf("\x1bc")
			fmt.Print("--- Welcome to Mixue ---\n\n")

			for _, menu := range HomeMenus {
				fmt.Printf("%d. %s\n", menu.ID, menu.Menu)
			}
			fmt.Print("\n0. Exit\n\n")

			fmt.Print("Choose a menu: ")

			choice, err := utils.InputInt(reader)

			if err != nil {
				panic("Invalid input, please enter a number...")
			}

			if choice == 0 {
				loop = false
				return
			}

			found := false
			for _, menu := range HomeMenus {
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
