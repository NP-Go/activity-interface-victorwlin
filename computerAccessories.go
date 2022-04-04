package main

import "fmt"

//Declare Computer Accessories Struct
type computerAccessories struct {
	title string
	price float64
}

//Create print method
func (ca computerAccessories) printPrice() {
	fmt.Printf("Title: %v  Price: $%.2f\n", ca.title, ca.price)
}