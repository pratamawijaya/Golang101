package main

import (
	"fmt"
)

func main() {
	var number uint8 = 255 // 0 - 255 
	var number1 uint16 = 65535 // 0 - 65535
	var number2 uint32 = 4294967295 // 0 - 4294967295 
	var number3 uint64 = 18446744073709551615 // 0 - 18446744073709551615
	var number4 uint = 0 // sama dengan uint32 atau uint64 (tergantung nilai)
	var number5 byte = 0 // sama dengan uint8
	var number6 int8 = -128 // -128 - 127
	var number7 int16 = -32768 // -32768 - 32767
	var number8 int32 = -2147483648 // -2147483648 - 2147483647
	var number9 int64 = -9223372036854775808
	var number10 int = -128 // sama dengan int32 atau int64 (tergantung nilai)
	var number11 rune = -2147483648 // sama dengan int32

	var decimalNumber = 2.62
	var exist bool = true


	fmt.Printf("bilangan : %d", number)
	fmt.Printf("bilangan : %d", number1)
	fmt.Printf("bilangan : %d", number2)
	fmt.Printf("bilangan : %d", number3)
	fmt.Printf("bilangan : %d", number4)
	fmt.Printf("bilangan : %d", number5)
	fmt.Printf("bilangan : %d", number6)
	fmt.Printf("bilangan : %d", number7)
	fmt.Printf("bilangan : %d", number8)
	fmt.Printf("bilangan : %d", number9)
	fmt.Printf("bilangan : %d", number10)
	fmt.Printf("bilangan : %d", number11)
	fmt.Printf("bilang decimal : %f", decimalNumber)
	
	if(exist){
		fmt.Printf("exist bool true")
	}else {
		fmt.Printf("bool false")
	}
		


}