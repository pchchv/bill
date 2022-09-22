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

func main() {}
