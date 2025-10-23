package main

import (
	"bufio"
	"fmt"
	"os"
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

		fmt.Println("1. Select Menu")
		fmt.Println("2. Cart")
		fmt.Println("3. History")

		fmt.Print("\n0. exit\n\n")

		fmt.Print("Choose a menu: ")
		input, _ := reader.ReadString('\n')
		menu := strings.TrimSpace(input)

		switch menu {
		case "1":
			fmt.Print("1")
		case "2":
			fmt.Print("2")
		case "3":
			fmt.Print("3")
		case "0":
			loop = false
		default:
			fmt.Print("Invalid menu option, press enter to continue...")
			scanner.Scan()
		}
	}
}
