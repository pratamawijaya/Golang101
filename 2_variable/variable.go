package main

import (
	"fmt"
)

func main() {
	// untuk declare variable, golang menggunaka keyword var 

	var firstName string = "John Snow"

	fmt.Println("Hello my name is " + firstName)
	fmt.Printf("\nHello my name is %s", firstName)
}