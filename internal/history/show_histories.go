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
		fmt.Println("\x1bc")
		fmt.Print("--------------- Your Histories ---------------\n\n")

		if len(models.Histories) == 0 {
			fmt.Print("Your histories is empty.\n\n")
			fmt.Print("----------------------------------------------\n\n")
			fmt.Print("Press Enter to go back to the main menu... ")
			scanner.Scan()
			loop = false
			continue
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

		fmt.Println("----------------------------------------------")
		fmt.Fprintln(w, "No\tNo. Invoice\tDate\tTotal")
		fmt.Fprintln(w, "---\t----------------\t--------\t-------------")

		for i, history := range models.Histories {
			fmt.Fprintf(w, "%d\t%s\t%s\tRp.%.2f\n", i+1, history.NoInvoice, history.Date, history.Total)
		}
		w.Flush()

		fmt.Print("----------------------------------------------\n\n")
		fmt.Print("0. Exit\n\n")
		fmt.Print("Enter number to view details: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)

		if err != nil {
			fmt.Print("Invalid input, please enter a number... ")
			scanner.Scan()
			continue
		}

		if choice == 0 {
			loop = false
			break
		}

		DetailsHistory(choice - 1)
	}
}
