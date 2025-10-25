package history

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/models"
	"os"
)

func DetailsHistory(index int) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\x1bc")

	if index < 0 || index >= len(models.Histories) {
		fmt.Println("History not found!")
		fmt.Print("Press Enter to continue...")
		scanner.Scan()
		return
	}

	history := models.Histories[index]

	fmt.Println("--- History Details ---")
	fmt.Println("Date:", history.Date)
	fmt.Println("No Invoice:", history.NoInvoice)
	fmt.Println("\nMenu Items:")
	for _, menu := range history.ListMenu {
		subtotal := float64(menu.Quantity) * menu.Price
		fmt.Printf("  - %s, Qty: %d, Price: Rp %.2f, Subtotal: Rp.%.2f\n", menu.Name, menu.Quantity, menu.Price, subtotal)
	}
	fmt.Printf("\nTotal: Rp %.2f\n", history.Total)
	fmt.Println("\n-----------------")

	fmt.Print("Press Enter to go back... ")
	scanner.Scan()
}
