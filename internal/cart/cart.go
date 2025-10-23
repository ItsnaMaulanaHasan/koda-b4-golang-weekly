package cart

type CartItem struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

var Carts []CartItem
