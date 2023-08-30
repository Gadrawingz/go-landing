package main

import "fmt"

func main() {

	age := 43
	// Down here, expression is omitted!
	switch {
	case age >= 0 && age <= 6:
		fmt.Printf("%d (Child) Your age greater than 0 and less than 6", age)
	case age >= 6 && age < 10:
		fmt.Printf("%d (Young) Your age greater than 6 and less than 10", age)
	case age >= 10 && age <= 19:
		fmt.Printf("%d (Teenager) Age is greater between 10 and 19", age)
	case age > 21 && age <= 39:
		fmt.Printf("%d (You are an adult and belong to the youth", age)
	case age > 39 && age <= 150:
		fmt.Printf("%d Too old, I hope you are enjoying your pension", age)
	case age > 151:
		fmt.Printf("%d age is greater than 100, You the zombie", age)
	default:
		fmt.Printf("Invalid age!")
	}
}
