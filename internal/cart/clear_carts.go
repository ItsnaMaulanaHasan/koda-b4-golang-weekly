package cart

import (
	"bufio"
	"context"
	"fmt"
	"golang-weekly/internal/utils"
	"os"
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
				conn, err := utils.GetConn()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
					os.Exit(1)
				}
				_, err = conn.Exec(context.Background(),
					`TRUNCATE TABLE carts`)
				if err != nil {
					panic(fmt.Sprintf("Unable to clear cart: %v", err))
				}
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
