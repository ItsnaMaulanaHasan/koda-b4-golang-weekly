package models

import (
	"bufio"
	"fmt"
	"text/tabwriter"
)

type Menu struct {
	ID    int
	Name  string
	Price float64
}

var Menus []Menu

type MenusPage struct {
	ID     int
	Menu   string
	Action func(reader *bufio.Reader, scanner *bufio.Scanner, w *tabwriter.Writer)
}

type Cart struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

var Carts = []Cart{}

type History struct {
	ID        int
	Date      string
	NoInvoice string
	ListMenu  []Cart
	Total     float64
}

var Histories = []History{}

func (menu *Menu) PrintOut(w *tabwriter.Writer) {
	for i, item := range Menus {
		fmt.Fprintf(w, "%d\t%s\tRp.%.2f\n", i+1, item.Name, item.Price)
	}
}

func (cart *Cart) PrintOut(w *tabwriter.Writer) {
	for i, item := range Carts {
		fmt.Fprintf(w, "%d\t%s\t%d\tRp.%.2f\n", i+1, item.Name, item.Quantity, item.Price*float64(item.Quantity))
	}
}

func (history *History) PrintOut(w *tabwriter.Writer) {
	for i, item := range Histories {
		fmt.Fprintf(w, "%d\t%s\t%s\tRp.%.2f\n", i+1, item.Date, item.NoInvoice, item.Total)
	}
}

type Printable interface {
	PrintOut(w *tabwriter.Writer)
}

func PrintRows(p Printable, w *tabwriter.Writer) {
	p.PrintOut(w)
}
