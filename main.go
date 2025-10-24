package main

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/menu"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer func() {
		fmt.Println("\nExiting program")
		os.Exit(0)
	}()
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	for loop {
		fmt.Printf("\x1bc")
		fmt.Print("--- Welcome to Mixue ---\n\n")

		for _, menu := range menu.MenuHomes {
			fmt.Printf("%d. %s\n", menu.ID, menu.Menu)
		}
		fmt.Print("\n0. exit\n\n")

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
		for _, menu := range menu.MenuHomes {
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
