package main

import "fmt"

func main() {
	/*
		Some more code to cover all variations of for loop
		--------------------------------------------------
		Let's print all even numbers from 0 to 40.
	*/

	/*
		In go, all the three components of the for loop namely initialisation,
		condition and post are optional.
		-> Initialisation and post are omitted.
		-> The semicolons in the for loop can also be omitted;
		-> This format can be considered as an alternative for while loop.
	*/

	// Recaps on values:

	a := 0
	for a <= 40 { // for ;a <= 10; {
		fmt.Printf("%d ", a)
		a += 2
	}

}
