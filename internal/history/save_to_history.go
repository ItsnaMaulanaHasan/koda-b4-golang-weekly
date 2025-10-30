package history

import (
	"context"
	"fmt"
	"golang-weekly/internal/models"
	"golang-weekly/internal/utils"
	"os"
)

func SaveToHistory(dataHistory *models.History) {
	conn, err := utils.GetConn()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	ctx := context.Background()

	var historyID int
	err = conn.QueryRow(ctx,
		`INSERT INTO histories (date, no_invoice, total) VALUES ($1, $2, $3) RETURNING id`,
		dataHistory.Date, dataHistory.NoInvoice, dataHistory.Total).Scan(&historyID)
	if err != nil {
		panic(fmt.Sprintf("Unable to save to histories: %v", err))
	}

	for i := range dataHistory.ListMenu {
		_, err = conn.Exec(ctx,
			`INSERT INTO product_history (history_id, product_id, quantity) VALUES ($1, $2, $3)`, historyID, dataHistory.ListMenu[i].Product_id, dataHistory.ListMenu[i].Quantity)
		if err != nil {
			panic(fmt.Sprintf("Unable to save to histories: %v", err))
		}
	}

	_, err = conn.Exec(context.Background(),
		`TRUNCATE TABLE carts`)
	if err != nil {
		panic(fmt.Sprintf("Unable to clear cart: %v", err))
	}
}
