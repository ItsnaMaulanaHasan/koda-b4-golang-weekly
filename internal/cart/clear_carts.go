package cart

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/models"
	"golang-weekly/internal/utils"
	"strings"
	"text/tabwriter"
)

func ClearCart(reader *bufio.Reader, scanner *bufio.Scanner, w *tabwriter.Writer) {
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
			printCarts(scanner, w)
			fmt.Print("Are you sure you want to clear carts (y/n)? ")
			choiceStr := utils.InputString(reader)
			if strings.ToLower(choiceStr) == "y" {
				models.Carts = []models.Cart{}
				fmt.Print("Carts successfully cleared! Press enter to continue... ")
				scanner.Scan()
				loop = false
				return
			} else if strings.ToLower(choiceStr) == "n" {
				loop = false
				return
			} else {
				panic("Invalid input, please enter 'y' or 'n'... ")
			}
		}()
	}
}
