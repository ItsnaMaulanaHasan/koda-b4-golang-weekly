package history

import (
	"golang-weekly/internal/data_menu"
)

type History struct {
	ID        int
	Date      string
	NoInvoice string
	ListMenu  []data_menu.MenuItem
	Total     int
}

var Histories = []History{}
