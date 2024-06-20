package main

import (
	"Phone/electronic"
	"fmt"
)

func printCharacteristics(p electronic.Phone) {
	fmt.Printf("Brand: %s, Model: %s, Type: %s\n", p.Brand(), p.Model(), p.Type())

	if sp, ok := p.(electronic.Smartphone); ok {
		fmt.Printf("OS: %s\n", sp.OS())
	}

	if stp, ok := p.(electronic.StationPhone); ok {
		fmt.Printf("Buttons Count: %d\n", stp.ButtonsCount())
	}
}

// Функция main
func main() {
	apple := electronic.NewApplePhone("iPhone 12")
	android := electronic.NewAndroidPhone("Samsung", "Galaxy S20")
	radio := electronic.NewRadioPhone("Panasonic", 15)

	printCharacteristics(apple)
	printCharacteristics(android)
	printCharacteristics(radio)
}
