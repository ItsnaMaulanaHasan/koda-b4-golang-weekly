package cart

import (
	"fmt"
	"golang-weekly/internal/models"
	"math/rand/v2"
	"time"
)

func getTotal(menuItems *[]models.CartItem) float64 {
	total := 0.0
	for _, item := range *menuItems {
		subtotal := float64(item.Quantity) * item.Price
		total += subtotal
	}
	return total
}

func createInvoice(menuItems *[]models.CartItem) models.History {
	ID := rand.IntN(9000) + 1000
	invoice := fmt.Sprintf("INV-MIXUE-%v", ID)

	dateNow := time.Now()
	dateStr := dateNow.Format("01-02-2006")

	dataHistory := models.History{
		ID:        ID,
		Date:      dateStr,
		NoInvoice: invoice,
		ListMenu:  *menuItems,
		Total:     getTotal(menuItems),
	}

	return dataHistory
}

func printInvoice(invoice *models.History) {
	fmt.Println("\x1bc")
	fmt.Println("--- INVOICE ---")
	fmt.Println("Date:", invoice.Date)
	fmt.Println("No Invoice:", invoice.NoInvoice)
	fmt.Println("\nMenu Items:")
	for _, menu := range invoice.ListMenu {
		subtotal := float64(menu.Quantity) * menu.Price
		fmt.Printf("  - %s, Qty: %d, Price: Rp %.2f, Subtotal: Rp.%.2f\n", menu.Name, menu.Quantity, menu.Price, subtotal)
	}
	fmt.Printf("\nTotal: Rp %.2f\n", invoice.Total)
	fmt.Println("\n-----------------")
}
