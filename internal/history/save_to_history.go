package history

import "golang-weekly/internal/models"

func SaveToHistory(dataHistory *models.History) {
	models.Histories = append(models.Histories, *dataHistory)
}
