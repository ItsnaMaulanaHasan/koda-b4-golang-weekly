package history

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/models"
	"golang-weekly/internal/utils"
	"text/tabwriter"
)

func ShowHistories(reader *bufio.Reader, scanner *bufio.Scanner, w *tabwriter.Writer) {
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

			if len(models.HistoryOrders.ListHistory) == 0 {
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

			models.PrintRows(models.HistoryOrders, w)

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

			if choice < 0 || choice > len(models.HistoryOrders.ListHistory) {
				panic("History not found!")
			} else {
				DetailsHistory(choice - 1)
			}

		}()
	}
}
