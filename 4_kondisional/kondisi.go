package main

import (
	"fmt"
)

func main() {
	var point = 10

	if point == 0 {
		fmt.Println("Poin Kosong")
	} else if point == 1 {
		fmt.Println("Poin 1")
	} else if point > 5 && point < 10 {
		fmt.Println("Lumayan")
	} else if point >= 10 {
		fmt.Println("Sempurna")
	}

	var nilai = 8840.0
	// percent is temporary variable
	if percent := nilai / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	var points = 6

	switch points {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}
