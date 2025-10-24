package datamenu

type MenuItem struct {
	ID    int
	Name  string
	Price float64
}

type MenusPage struct {
	ID     int
	Menu   string
	Action func()
}
