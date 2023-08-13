package main

import (
	"fmt"
)

func main() {

	// Exploring go constants
	const birthYear = 1990
	fmt.Println("My birth year is: ", birthYear)

	// Declaring a group of constants
	const (
		firstName    = "Joyner"
		lastName     = "Lucas"
		personAge    = 40
		personHeight = 1.80
		isSingle     = true
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(personAge)
	fmt.Println(personHeight)
	fmt.Println(isSingle)

	// String Constants
	const tradeName = "Software Development"
	fmt.Println(tradeName)

	// Being Strongly typed language
	const wifeName = "Mariah Carey"
	var realWifeName = wifeName
	fmt.Printf("Type is: %T, then value %v \n", wifeName, wifeName)
	fmt.Printf("Type is: %T, then value %v \n", realWifeName, realWifeName)

	// Creating a typed constant
	const typedName string = "Clara"
	const typedAge int = 80
	fmt.Println("*********************************")
	fmt.Println("The typed name is : ", typedName)
	fmt.Println("The typed age is : ", typedAge)

	// Mixing typed and untyped
	var mainSchool = "Cambridge" // Allowed
	type myString string
	var customSchool myString = "Liverpool University" // Allowed
	fmt.Println(customSchool, ", ", mainSchool)

	// Boolean Constants
	const isSingleForever = true
	type myBoolean bool
	var defaultBoolean = isSingleForever          // Allowed
	var customBoolean myBoolean = isSingleForever // Allowed
	fmt.Println(isSingleForever)
	fmt.Println(defaultBoolean) // Is always true haha
	fmt.Println(customBoolean)

	// Numeric Constants
	const numberOne int = 1
	var numberOneIntVar int = numberOne
	fmt.Println(numberOneIntVar)
	//var numberOneInt32Var int32 = numberOne
	//var numberOneFloat64Var float64 = numberOne
	//var numberOneComplex64var complex64 = numberOne
	//fmt.Println("intVar", numberOne, "\nint32Var", numberOneIntVar, "\nfloat64Var", numberOneInt32Var, "\ncomplex64Var", numberOneFloat64Var, ": ", numberOneComplex64var)

}
