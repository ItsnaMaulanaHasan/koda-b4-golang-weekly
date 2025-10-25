package history

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/models"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func ShowHistories() {
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)

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

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

			fmt.Println("----------------------------------------------")
			fmt.Fprintln(w, "No\tDate\tNo. Invoice\tTotal")
			fmt.Fprintln(w, "---\t----------\t--------------\t-------------")

			models.PrintRows(models.HistoryOrders, w)

			w.Flush()

			fmt.Print("----------------------------------------------\n\n")
			fmt.Print("0. Exit\n\n")
			fmt.Print("Enter number to view details: ")

			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)
			choice, err := strconv.Atoi(choiceStr)

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
