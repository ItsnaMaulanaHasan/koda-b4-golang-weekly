package models

import (
	"bufio"
	"fmt"
	"text/tabwriter"
)

type MenusPage struct {
	ID     int
	Menu   string
	Action func(reader *bufio.Reader, scanner *bufio.Scanner, w *tabwriter.Writer)
}
type MenuItem struct {
	ID    int
	Name  string
	Price float64
}

type Menus struct {
	ListMenu []MenuItem
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

func (menu *Menus) PrintOut(w *tabwriter.Writer) {
	for i, item := range menu.ListMenu {
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

var MenuMixue = &Menus{ListMenu: []MenuItem{
	{ID: 1, Name: "Mixue Ice Cream", Price: 8000},
	{ID: 2, Name: "BOBA Sundae", Price: 16000},
	{ID: 3, Name: "Strawberry Mi-Shake", Price: 16000},
	{ID: 4, Name: "BOBA Mi-Shake", Price: 16000},
	{ID: 5, Name: "Chocolate Cookies Smoothies", Price: 16000},
	{ID: 6, Name: "Brown Sugar Pearl Milk Tea", Price: 19000},
	{ID: 7, Name: "Pearl Milk Tea", Price: 22000},
	{ID: 8, Name: "Oats Milk Tea", Price: 22000},
	{ID: 9, Name: "Coconut Jelly Milk Tea", Price: 22000},
	{ID: 10, Name: "Red Bean Milk Tea", Price: 22000},
	{ID: 11, Name: "Fresh Squeezed Lemonade", Price: 10000},
	{ID: 12, Name: "Peach Earl Grey Tea", Price: 16000},
	{ID: 13, Name: "Original Jasmine Tea", Price: 10000},
	{ID: 14, Name: "Original Earl Grey Tea", Price: 10000},
}}

var CartOrders = &Carts{}

var HistoryOrders = &Histories{}
