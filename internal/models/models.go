package models

import (
	"fmt"
	"text/tabwriter"
)

type MenuItem struct {
	ID    int
	Name  string
	Price float64
}

type Menus struct {
	ListMenu []MenuItem
}

func (menu Menus) PrintOut() []string {
	var results []string
	for i, item := range menu.ListMenu {
		results = append(results, fmt.Sprintf("%d\t%s\tRp.%.2f\n", i+1, item.Name, item.Price))
	}
	return results
}

type MenusPage struct {
	ID     int
	Menu   string
	Action func()
}

type CartItem struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

type History struct {
	ID        int
	Date      string
	NoInvoice string
	ListMenu  []CartItem
	Total     float64
}

type Printable interface {
	PrintOut() []string
}

func PrintRows(p Printable, w *tabwriter.Writer) {
	for _, line := range p.PrintOut() {
		fmt.Fprint(w, line)
	}
}

var MenuMixue = Menus{ListMenu: []MenuItem{
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

var Carts = []CartItem{}

var Histories = []History{}
