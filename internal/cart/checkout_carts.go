package cart

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/models"
	"strings"
)

func CheckoutCart(reader *bufio.Reader, scanner *bufio.Scanner) {
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
			fmt.Print("Are you sure you want to checkout (y/n)? ")
			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)
			if strings.ToLower(choiceStr) == "y" {
				CreateInvoice(models.CartOrders.ListCart)
				models.CartOrders.ListCart = []models.CartItem{}
				fmt.Print("Checkout successful! Press enter to continue... ")
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
