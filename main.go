package main

type Bill struct {
	list          []Item
	totalQuantity int
	totalSum      float64
}

type Item struct {
	name     string
	price    float64
	quantity int
}

func main() {}
