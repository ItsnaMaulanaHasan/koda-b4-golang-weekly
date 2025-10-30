package main

import (
	"bufio"
	"context"
	"fmt"
	"golang-weekly/internal/cart"
	"golang-weekly/internal/history"
	"golang-weekly/internal/menu"
	"golang-weekly/internal/models"
	"golang-weekly/internal/utils"
	"os"
	"text/tabwriter"

	"github.com/jackc/pgx/v5"
)

var HomeMenus = []models.MenusPage{
	{ID: 1, Menu: "Select Menu", Action: menu.SelectMenu},
	{ID: 2, Menu: "Cart", Action: cart.ShowCarts},
	{ID: 3, Menu: "History", Action: history.ShowHistories},
	// {ID: 4, Menu: "Clear Cache", Action: utils.ClearCache},
	{ID: 4, Menu: "Clear Cache", Action: nil},
}

func main() {
	defer func() {
		fmt.Printf("\x1bc")
		fmt.Println("\nExiting program")
		os.Exit(0)
	}()

	conn, err := utils.GetConn()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, _ := conn.Query(context.Background(), "SELECT id, name, price FROM products")
	models.Menus, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Menu])
	if err != nil {
		panic(err)
	}

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
