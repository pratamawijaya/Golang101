package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slice")

	var fruits = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits[0:2]

	fmt.Println(fruits[0])
	fmt.Println(newFruits) // ["apple", "grape"]
	fmt.Println(len(fruits))
	var cFruits = append(fruits, "papaya")

	fmt.Println(cFruits)

	// slice merupakan tipe data reference

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]
}
