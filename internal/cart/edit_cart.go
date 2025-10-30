package cart

import (
	"bufio"
	"context"
	"fmt"
	"golang-weekly/internal/models"
	"golang-weekly/internal/utils"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func EditCart(reader *bufio.Reader, scanner *bufio.Scanner, w *tabwriter.Writer) {
	conn, err := utils.GetConn()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
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

			loop = printCarts(scanner, w)
			if !loop {
				return
			}

			fmt.Print("0. Exit\n")

			fmt.Print("\nEnter the menu number you want to edit: ")

			choice, err := utils.InputInt(reader)

			if err != nil {
				panic("Invalid input, please enter a number... ")
			}

			if choice == 0 {
				loop = false
				return
			}

			found := false
			for _, item := range models.Carts {
				if item.Product_id == choice {
					fmt.Printf("Enter new quantity for %s: ", item.Name)
					quantityStr, _ := reader.ReadString('\n')
					quantityStr = strings.TrimSpace(quantityStr)
					quantity, err := strconv.Atoi(quantityStr)
					if err != nil || quantity < 0 {
						panic("Invalid quantity, please enter a valid number... ")
					} else if quantity == 0 {
						_, err := conn.Exec(context.Background(),
							`DELETE FROM carts WHERE id = $1`,
							item.ID)
						if err != nil {
							panic(fmt.Sprintf("Unable to update cart: %v", err))
						}
						fmt.Printf("%s removed from cart. Press enter to continue... ", item.Name)
					} else {
						_, err := conn.Exec(context.Background(),
							`UPDATE carts SET quantity = $1 WHERE id = $2`, quantity,
							item.ID)
						if err != nil {
							panic(fmt.Sprintf("Unable to update cart: %v", err))
						}
						fmt.Printf("Quantity for %s updated to %d. Press enter to continue... ", item.Name, quantity)
					}
					scanner.Scan()
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
