package main

import (
	"fmt"
	"math"
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

	// Why the value of a constant should be known at compile time
	var squaredNumber1 = math.Sqrt(5) // Allowed
	//const squaredNumber2 = math.Sqrt(5) // Not Allowed
	fmt.Println(squaredNumber1, "Allowed")

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
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
	fmt.Printf("The type of intVar is : %T, float64var: %T, complex64Var : %T \n", intVar, float64Var, complex64Var)

	/*
		NB1: Numeric constants are free to be mixed and matched in expressions and
		a type is needed only when they are assigned to variables or used in any
		place in code which demands a type.
	*/
	var mixedExpression = 7.82 / 4
	fmt.Println("The value is:", mixedExpression)
	fmt.Printf("The type: %T, The value: %v", mixedExpression, mixedExpression)

}
