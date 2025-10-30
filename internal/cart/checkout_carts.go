package cart

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/history"
	"golang-weekly/internal/models"
	"golang-weekly/internal/utils"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

func CheckoutCart(reader *bufio.Reader, scanner *bufio.Scanner, w *tabwriter.Writer) {
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
			fmt.Print("Are you sure you want to checkout (y/n)? ")
			choiceStr := utils.InputString(reader)
			if strings.ToLower(choiceStr) == "y" {
				invoice := make(chan models.History, 1)

				var wg sync.WaitGroup

				wg.Add(1)
				go func() {
					defer wg.Done()
					fmt.Print("Create invoice... ")
					invoice <- createInvoice(&models.Carts)
					time.Sleep(2 * time.Second)
					fmt.Println("✅")
				}()
				wg.Wait()

				inv := <-invoice

				wg.Add(1)
				go func() {
					defer wg.Done()
					fmt.Print("Saving to history... ")
					history.SaveToHistory(&inv)
					time.Sleep(3 * time.Second)
					fmt.Println("✅")
				}()
				wg.Wait()

				wg.Add(1)
				go func() {
					defer wg.Done()
					fmt.Print("Printing invoice... ")
					time.Sleep(4 * time.Second)
					printInvoice(&inv)
				}()
				wg.Wait()

				time.Sleep(200 * time.Millisecond)

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
