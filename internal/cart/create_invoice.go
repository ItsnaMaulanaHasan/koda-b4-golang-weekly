package cart

import (
	"fmt"
	"golang-weekly/internal/models"
	"math/rand/v2"
	"time"
)

func getTotal(menuItems []models.CartItem) float64 {
	total := 0.0
	for _, item := range menuItems {
		subtotal := float64(item.Quantity) * item.Price
		total += subtotal
	}
	return total
}

func CreateInvoice(menuItems []models.CartItem) {
	ID := rand.IntN(9000) + 1000
	invoice := fmt.Sprintf("INV-MIXUE-%v", ID)

	dateNow := time.Now()
	dateStr := dateNow.Format("01-02-2006")

	models.HistoryOrders.ListHistory = append(models.HistoryOrders.ListHistory, models.History{
		ID:        ID,
		Date:      dateStr,
		NoInvoice: invoice,
		ListMenu:  menuItems,
		Total:     getTotal(menuItems),
	})
}
