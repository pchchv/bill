package main

type Bill struct {
	list          []Item  // List of ordered items
	totalQuantity int     // Total number of items ordered
	totalSum      float64 // Total order amount
}

type Item struct {
	name     string
	price    float64
	quantity int
}

// List of ordered items for order picking
type KitchenList struct {
	list          []Item
	totalQuantity int
}

func main() {}
