package models

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang-weekly/internal/utils"
	"text/tabwriter"
)

type Menu struct {
	ID    int
	Name  string
	Price float64
}

func getDataMenu() []Menu {
	var menus []Menu
	dataMenu := utils.GetData("https://raw.githubusercontent.com/ItsnaMaulanaHasan/koda-b4-golang-weekly-data/refs/heads/main/data.json")

	err := json.Unmarshal(dataMenu, &menus)

	if err != nil {
		panic("Failed to get data menu")
	}

	return menus
}

var Menus = getDataMenu()

type MenusPage struct {
	ID     int
	Menu   string
	Action func(reader *bufio.Reader, scanner *bufio.Scanner, w *tabwriter.Writer)
}

type CartItem struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

type Carts struct {
	ListCart []CartItem
}

type History struct {
	ID        int
	Date      string
	NoInvoice string
	ListMenu  []CartItem
	Total     float64
}

type Histories struct {
	ListHistory []History
}

func (menu *Menu) PrintOut(w *tabwriter.Writer) {
	for i, item := range Menus {
		fmt.Fprintf(w, "%d\t%s\tRp.%.2f\n", i+1, item.Name, item.Price)
	}
}

func (cart *Carts) PrintOut(w *tabwriter.Writer) {
	for i, item := range cart.ListCart {
		fmt.Fprintf(w, "%d\t%s\t%d\tRp.%.2f\n", i+1, item.Name, item.Quantity, item.Price*float64(item.Quantity))
	}
}

func (history *Histories) PrintOut(w *tabwriter.Writer) {
	for i, item := range history.ListHistory {
		fmt.Fprintf(w, "%d\t%s\t%s\tRp.%.2f\n", i+1, item.Date, item.NoInvoice, item.Total)
	}
}

type Printable interface {
	PrintOut(w *tabwriter.Writer)
}

func PrintRows(p Printable, w *tabwriter.Writer) {
	p.PrintOut(w)
}

var CartOrders = &Carts{}

var HistoryOrders = &Histories{}
