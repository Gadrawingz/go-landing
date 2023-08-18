package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(":")

	// Go Booleans
	isSerious := false
	dataFound := true
	truthy := true
	falsy := false
	result1 := isSerious && dataFound
	result2 := isSerious || dataFound
	result3 := truthy || truthy
	result4 := falsy || falsy
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Signed Integers
	var numberOne int8 = 127 // range: -128 to 127
	var numberXBack int8 = -111
	var numberTwo int16 = 4554              // range: -32768 to 32767
	var myInt32 int32 = 400343222           // range: -2147483648 to 2147483647
	var myInt64 int64 = 9223372036854775805 // range: -2147483648 to 2147483647
	var numberThree rune = 400              // range:
	fmt.Println(numberOne)
	fmt.Println(numberTwo)
	fmt.Println(numberThree)
	fmt.Println(numberXBack, " ", myInt32, " ", myInt64)

	// int represents 32/64 bit integers depending on the underlying platform
	// Using unsafe:
	var grandMaAge int = 101
	grandPaAge := 97
	fmt.Println("GrandMa age Value: ", grandMaAge, " The type is: ", grandMaAge)
	fmt.Println("GrandPa age Value: ", grandPaAge, " The type is: ", grandPaAge)
	fmt.Printf("GrandPa age Type is: %T, The Size is: %d \n", grandPaAge, unsafe.Sizeof(grandPaAge))
	fmt.Printf("GrandMa age Type is: %T, The Size is: %d \n", grandMaAge, unsafe.Sizeof(grandMaAge))

	// Unsigned Integers
	var x uint8 = 40
	var y uint64 = 322
	fmt.Println(x)
	fmt.Println(y)

	// Floating Points
	var weeklyWage float64 = 4000
	var monthlyWage float32 = 3000
	hisSalary, herSalary := 33000.5, 200.5
	salaryTotal := hisSalary + herSalary
	salaryDifference := hisSalary - herSalary
	fmt.Println(weeklyWage)
	fmt.Println(monthlyWage)
	fmt.Println(hisSalary)
	fmt.Println(herSalary)
	fmt.Println("Total salary is: ", salaryTotal)
	fmt.Println("Difference in salary is: ", salaryDifference)

	// Complex Types (Complex numbers):
	// Using shorthand syntax
	myComplexNumber := 4 + 2i
	fmt.Println("My Complex Number is: ", myComplexNumber)

	// Understanding Complex Numbers
	comp1 := complex(5, 7)
	comp2 := 8 + 13i
	compAddition := comp1 + comp2
	fmt.Println("Sum equal to: ", compAddition) // (13+20i)
	compMultiply := comp1 * comp2
	fmt.Println("Multiplication: ", compMultiply) // (-51+121i)

	// String Types
	playerFirstName := "James"
	playerPosition := "Captain"
	playerLastName := "Milner"
	fmt.Println("Full Name is: ", playerFirstName+" "+playerLastName, ", He's ", playerPosition)
	fmt.Println("Full Name is: ", playerFirstName, playerLastName)

	/*
		Other types:
		byte: is an alias of uint8
		rune: is an alias of int32
	*/
	var newBalance byte = 150        // Same as uint8, 8-bits (range 0-255)
	var childSaving rune = 535554400 // Same as uint32, 32-bits (range 0-42949672955)
	fmt.Println(newBalance)
	fmt.Println(childSaving)

	// Go Type Conversion
	numberA := 32.4
	numberB := 332
	numberC := 75.52
	numberD := 112

	// The same with assignment case
	var newlyAssigned = int(numberC)

	// int + float64 + bool not allowed
	sumOfNumbers := numberB + int(numberC)
	fmt.Println("Sum 01: ", sumOfNumbers)             // float to int
	fmt.Println("Sum 02: ", numberA+float64(numberD)) // int to float
	fmt.Println("Newly assigned variable is: ", newlyAssigned)

}
