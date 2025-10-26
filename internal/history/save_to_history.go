package history

import "golang-weekly/internal/models"

func SaveToHistory(dataHistory *models.History) {
	models.HistoryOrders.ListHistory = append(models.HistoryOrders.ListHistory, *dataHistory)
}
