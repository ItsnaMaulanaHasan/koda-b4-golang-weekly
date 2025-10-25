package main

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/cart"
	"golang-weekly/internal/history"
	"golang-weekly/internal/menu"
	"golang-weekly/internal/models"
	"os"
	"strconv"
	"strings"
)

var HomeMenus = []models.MenusPage{
	{ID: 1, Menu: "Select Menu", Action: menu.SelectMenu},
	{ID: 2, Menu: "Cart", Action: cart.CartsPage},
	{ID: 3, Menu: "History", Action: history.ShowHistories},
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
	for loop {
		fmt.Printf("\x1bc")
		fmt.Print("--- Welcome to Mixue ---\n\n")

		for _, menu := range HomeMenus {
			fmt.Printf("%d. %s\n", menu.ID, menu.Menu)
		}
		fmt.Print("\n0. Exit\n\n")

		fmt.Print("Choose a menu: ")

		input, _ := reader.ReadString('\n')
		choiceStr := strings.TrimSpace(input)
		choice, err := strconv.Atoi(choiceStr)

		if err != nil {
			fmt.Print("Invalid input, please enter a number...")
			scanner.Scan()
			continue
		}

		if choice == 0 {
			loop = false
			break
		}

		found := false
		for _, menu := range HomeMenus {
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
