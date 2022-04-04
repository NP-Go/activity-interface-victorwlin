package main

//Declare Interface
type product interface {
	printPrice()
}

type list []product

//create methods for print
func (list list) printAll() {
	for _, value := range list {
		value.printPrice()
	}
}