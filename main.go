package main

import (
	"fmt"
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

func creator(order *[]Item) ([]byte, []byte) {
	var bill Bill
	var kitchenList KitchenList
	var tq int
	var ts float64
	bill.list = *order
	for _, i := range *order {
		tq += i.quantity
		ts += i.price * float64(i.quantity)
	}
	bill.totalQuantity = tq
	bill.totalSum = ts
	kitchenList.list = *order
	kitchenList.totalQuantity = tq
	return billFormater(bill), listFormater(kitchenList)
}

func billFormater(bill Bill) []byte {
	fb := fmt.Sprintf("%25v\n\n\n", getEnvValue("NAME"))
	for _, item := range bill.list {
		for i := 0; i < item.quantity; i++ {
			fb += fmt.Sprintf("%-25v ...$%v\n", item.name+":", item.price)
		}
	}
	fb += fmt.Sprintf("\n%-25v ...$%0.2f", "total:", bill.totalSum)
	return []byte(fb)
}

func listFormater(list KitchenList) []byte {
	fl := "\n"
	for _, item := range list.list {
		fl += fmt.Sprintf("%-25v ...%v\n", item.name+":", item.quantity)

	}
	fl += fmt.Sprintf("\n%-25v ...%v", "total:", list.totalQuantity)
	return []byte(fl)
}

func main() {
}
