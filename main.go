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
		choice := strings.TrimSpace(input)

		if choice == "0" {
			loop = false
			continue
		}

		found := false
		for _, menu := range menu.MenuHomes {
			if strconv.Itoa(menu.ID) == choice {
				if menu.Action != nil {
					menu.Action()
				} else {
					fmt.Println("Menu action not implemented yet.")
				}
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Invalid menu option, press enter to continue...")
			scanner.Scan()
		}
	}
}
