package main

// import fmt https://golang.org/pkg/fmt/
import (
	"fmt"
	"math"
)

func main() {
	var firstName string = "Pratama"
	var lastName string
	// penulisan variable dengan type inference
	hoby := "Sepeda"

	lastName = "Wijaya"

	fmt.Printf("Hello World %s %s %s \n", firstName, lastName, hoby)

	fmt.Println("Hello", firstName, lastName+" !", hoby)

	// konstanta
	const phi = math.Pi

	fmt.Printf("bilangan konstant phi %f\n", phi)

	// type data
	var positiveNumber uint8 = 99
	var negativeNumber = -1234

	fmt.Printf("bilangan positive %d\n", positiveNumber)
	fmt.Printf("bilangan negative %d\n", negativeNumber)

	// kondisional

	var point = 4

	if point == 10 {
		fmt.Println("lulus dengan sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Println("tidak lulus, nilai anda ", point)
	}

	var nilai = 8840.0

	if percent := nilai / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	var check = 6

	switch check {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[2:4]

	fmt.Println("panjang ", len(slice1))
	fmt.Println("capacity ", cap(slice1))

}
