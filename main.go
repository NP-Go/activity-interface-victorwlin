package main

func main() {
	//Declare objects
	minecraft := game{"Minecraft", 5.00}
	wow := game{"World of Warcraft", 19.00}
	elite := game{"Elite Dangerous", 54.00}
	candle := book{"Candle in the Tomb", 20.00}
	barney := book{"Barney and Friends", 10.00}
	razerEar := computerAccessories{"Razer BT Earpiece", 159.00}
	razerKeyboard := computerAccessories{"Razer Keyboard", 110.00}
	logiMouse := computerAccessories{"Logitech Mouse", 80.00}

	//Insert code here
	var inventory list
	inventory = append(inventory, minecraft, wow, elite, candle, barney, razerEar, razerKeyboard, logiMouse)
	inventory.printAll()
}
