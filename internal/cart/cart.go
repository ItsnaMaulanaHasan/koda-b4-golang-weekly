package cart

import "golang-weekly/internal/datamenu"

type CartItem struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

var Carts []CartItem

var CartMenus = []datamenu.MenusPage{
	{ID: 1, Menu: "Checkout", Action: nil},
	{ID: 2, Menu: "Edit Cart", Action: nil},
	{ID: 3, Menu: "Clear Cart", Action: nil},
}
