package main

// import fmt https://golang.org/pkg/fmt/
import "fmt"

func main()  {
	var firstName string = "Pratama"
	var lastName string

	lastName = "Wijaya"

	fmt.Printf("Hello World %s %s \n", firstName, lastName)

	fmt.Println("Hello", firstName, lastName + " !")

}