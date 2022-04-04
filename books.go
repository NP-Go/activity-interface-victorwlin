package main

import "fmt"

//Declare books Struct
type book struct {
	title string
	price float64
}

//Create print method
func (b book) printPrice() {
	fmt.Printf("Title: %v  Price: $%.2f\n", b.title, b.price)
}