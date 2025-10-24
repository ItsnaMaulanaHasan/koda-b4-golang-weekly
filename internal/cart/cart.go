package cart

import "golang-weekly/internal/data_menu"

type CartItem struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

var Carts []CartItem

var CartMenus = []data_menu.MenusPage{
	{ID: 1, Menu: "Checkout", Action: CheckoutCart},
	{ID: 2, Menu: "Edit Cart", Action: EditCart},
	{ID: 3, Menu: "Clear Cart", Action: ClearCart},
}
