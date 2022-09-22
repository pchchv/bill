package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

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

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value. Outputs a panic if the value is missing
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

func creator(order *[]Item) (Bill, KitchenList) {
	var bill Bill
	var kitchenList KitchenList
	return bill, kitchenList
}

func main() {}
