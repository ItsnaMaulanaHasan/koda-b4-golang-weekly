package history

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowHistories() {
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	for loop {
		fmt.Println("\x1bc")
		fmt.Print("--- Your Histories ---\n\n")

		if len(Histories) == 0 {
			fmt.Print("Your cart is empty.\n\n")
			fmt.Print("-----------------\n\n")
			fmt.Print("Enter to go back to the main menu... ")
			scanner.Scan()
			loop = false
			continue
		}

		for i, history := range Histories {
			fmt.Printf("%d. ID: %d - Date: %s - No Invoice: %s - Total: Rp %d\n", i+1, history.ID, history.Date, history.NoInvoice, history.Total)
		}
		fmt.Print("\n-----------------\n\n")
		fmt.Print("\n0. exit\n")
		fmt.Println("Enter number to view details: ")

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

		found := false
		for _, history := range Histories {
			if choice == history.ID {
				fmt.Println("=== History Details ===")
				fmt.Println("ID:", history.ID)
				fmt.Println("Date:", history.Date)
				fmt.Println("No Invoice:", history.NoInvoice)
				fmt.Println("Menu Items:")
				for _, menu := range history.ListMenu {
					fmt.Println("- ", menu.Name, "Price:", menu.Price)
				}
				fmt.Println("Total:", history.Total)
				found = true
				break
			}
		}
		if !found {
			fmt.Println("History ID not found, press enter to continue...")
			scanner.Scan()
		}
	}
}
