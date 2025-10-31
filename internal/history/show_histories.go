package history

import (
	"bufio"
	"context"
	"fmt"
	"golang-weekly/internal/models"
	"golang-weekly/internal/utils"
	"os"
	"text/tabwriter"

	"github.com/jackc/pgx/v5"
)

func ShowHistories(reader *bufio.Reader, scanner *bufio.Scanner, w *tabwriter.Writer) {
	conn, err := utils.GetConn()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	ctx := context.Background()

	rows, err := conn.Query(ctx, "SELECT id, date, no_invoice, total FROM histories ORDER BY date DESC")
	if err != nil {
		panic(fmt.Sprintf("Query histories failed: %v", err))
	}

	histories, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.History])
	rows.Close()
	if err != nil {
		panic(fmt.Sprintf("Failed to collect histories: %v", err))
	}

	fmt.Println(histories)

	for i := range histories {
		detailRows, err := conn.Query(ctx, `
			SELECT
				product_history.id,
				product_history.product_id,
				products.name,
				product_history.quantity,
				products.price
			FROM product_history
			JOIN products ON products.id = product_history.product_id
			WHERE product_history.history_id = $1
		`, histories[i].ID)

		if err != nil {
			panic(fmt.Sprintf("Query product_history failed: %v", err))
		}

		carts, err := pgx.CollectRows(detailRows, pgx.RowToStructByName[models.Cart])
		detailRows.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to collect carts: %v", err))
		}

		histories[i].ListMenu = carts
	}

	models.Histories = histories
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
			fmt.Print("--------------- Your Histories ---------------\n\n")

			if len(models.Histories) == 0 {
				fmt.Print("Your histories is empty.\n\n")
				fmt.Print("----------------------------------------------\n\n")
				fmt.Print("Press Enter to go back to the main menu... ")
				scanner.Scan()
				loop = false
				return
			}

			fmt.Println("----------------------------------------------")
			fmt.Fprintln(w, "No\tDate\tNo. Invoice\tTotal")
			fmt.Fprintln(w, "---\t----------\t--------------\t-------------")

			models.PrintRows(&models.History{}, w)

			w.Flush()

			fmt.Print("----------------------------------------------\n\n")
			fmt.Print("0. Exit\n\n")
			fmt.Print("Enter number to view details: ")

			choice, err := utils.InputInt(reader)

			if err != nil {
				panic("Invalid input, please enter a number... ")
			}

			if choice == 0 {
				loop = false
				return
			}

			if choice < 0 || choice > len(models.Histories) {
				panic("History not found!")
			} else {
				DetailsHistory(choice-1, scanner)
			}

		}()
	}
}
