package main

// import fmt https://golang.org/pkg/fmt/
import "fmt"
import "math"

func main()  {
	var firstName string = "Pratama"
	var lastName string
	// penulisan variable dengan type inference
	hoby := "Sepeda"

	lastName = "Wijaya"

	fmt.Printf("Hello World %s %s %s \n", firstName, lastName, hoby)

	fmt.Println("Hello", firstName, lastName + " !", hoby)

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

}
