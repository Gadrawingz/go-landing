package main

import "fmt"

func main() {
	// Variable learning @gadiradufasha!
	// (1) Single variable declaration

	// Variable declaration with initial value:
	// var name type = initial-value
	var currentYear int = 2023
	fmt.Println("The current year is ", currentYear)

	// Variable declaration where it can be re-assigned!
	var randomYear int
	randomYear = 2003 // Re-assignment 1
	fmt.Println("The random year is ", randomYear)
	randomYear = 2004 // Re-assignment 2
	fmt.Println("The random year is ", randomYear)

	// Type inference
	// We can see that the type int of the variable age has been removed.
	// Since the variable has an initial value 4400, Go can infer it is of type int.

	var totalAmount = 4400
	var numberOfApples = 45
	// Above types will be inferred

	fmt.Println(totalAmount)
	fmt.Println(numberOfApples)

	// Multiple variables declaration
	var boyAge, girlAge int = 20, 18
	var minHeight, maxHeight, minWidth, maxWidth = 300, 80, 500, 400 // "int" is dropped
	var oddNumber, evenNumber float64 = 65.70, 10.44
	var hasBoyfriend, goodEnough, isWrong bool = true, false, true
	var firstName, middleName, lastName string = "Lucky", "Phillips", "Dube"
	fmt.Println(boyAge, girlAge)
	fmt.Println(maxHeight, minWidth, maxWidth, minHeight)
	fmt.Println(oddNumber, evenNumber)
	fmt.Println(hasBoyfriend, goodEnough, isWrong)
	fmt.Println("Is she has boyfriend? ", hasBoyfriend, " is it good enough?", goodEnough)
	fmt.Println("The full name is ", firstName, " ", middleName, " ", lastName, "!")

	// INITIAL, DEFAULT VALUES:
	// If the initial value is not specified for variables,
	// They will have 0 assigned as their initial value.
	var catLegs, dogLegs int
	fmt.Println("Cat Legs: ", catLegs, "Dog Legs: ", dogLegs)
	catLegs, dogLegs = 4, 5
	fmt.Println("Cat Legs: ", catLegs, "Dog Legs: ", dogLegs)

	// The cases where we would want to declare variables belonging to
	// different types in a single statement
	var (
		studentName        = "Lil Wayne"
		studentAge         = 30
		studentHeight      = 1.77
		studentBodyCount   int
		studentIndexNumber = 43432344
	)

	// Display 'em
	fmt.Println(studentName)
	fmt.Println(studentAge)
	fmt.Println(studentHeight)
	fmt.Println(studentBodyCount)
	fmt.Println(studentIndexNumber)
	
	// Shorthand for writing variables
	girlsNumber := 50
	boysNumber := 45
	teachersNumber := 19
	// Here above, Go automatically infers that girlsNumber, boysNumber is of type int
	var totalChildren int
	totalChildren = boysNumber + girlsNumber
	fmt.Println("The result of all children is: ", totalChildren)
	fmt.Println("All teachers we have are : ", teachersNumber)

	// Shorthand notation for declaring multiple variables in a single line
	initial := "His"

	// Declaration of 3 variables of type string and int respectively
	childFirstName, childNickName, childAge := "Gad", "Iradufasha", 20
	fmt.Println(initial, " name is ", childFirstName, " ", childNickName, " and ", childAge, " years old!")

	// 4: Shorthand NB:
	// This shorthand can only be used when at least on of the variable
	// On the left side of := is newly declared

	number1, number2 := 599, 377
	fmt.Println("Number 1: ", number1, " Number 2: ", number2)

	// Number2 is already declared but Number2 is new
	number2, number3 := 328, 260
	fmt.Println("Number 2: ", number2, " Number 3: ", number3)

	// Assigning new values to already declared variables
	number1, number2, childAge = 12, 44, 3

	/*
		Using this shorthand when there is no new variable result to error!
		number1, number2, childAge := 12, 44, 3
	*/
	fmt.Println("Number 1: ", number1, " Number 2: ", number2, " Child Age : ", childAge)

	// 2: Assigning variable values which are computed during run time is possible
	n1, n2 := 394.4, 64.5
	minimumValue := math.Min(n1, n2)
	maximumValue := math.Max(n2, n1)
	fmt.Println("Maximum value is: ", maximumValue)
	fmt.Println("Minimum value is: ", minimumValue)

	// Go being strongly typed.
	// Variables declared as belonging to one type cannot be assigned a value of another type.
	sumOfStudents := 4000
	// sumOfStudents = "Gad"
	fmt.Println("Total amount is : ", sumOfStudents)
}
