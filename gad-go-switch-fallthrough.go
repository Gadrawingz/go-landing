package main

import "fmt"

func myNumber() int {
	num := 20 * 5
	return num
}

func main() {
	/*
		This go program will check whether the input number is less than
		50, 100, or 200 using myNumber function!
	*/
	fmt.Println("The returned number is: ", myNumber())

	// Let us use switch case:
	// n is not a constant!
	switch n := myNumber(); {

	case n < 45:
		fmt.Printf("%d is less than 45 \n", n)
		fallthrough
		
	case n < 81:
		fmt.Printf("%d is less than 90 \n", n)
		fallthrough
		
	case n < 98:
		fmt.Printf("%d is less than 90 \n", n)
		fallthrough
		
	case n < 120:
		fmt.Printf("%d is less than 120 \n", n)
		fallthrough
		
	case n < 200:
		fmt.Printf("%d is less than 200 \n", n)
		fallthrough
		
	case n < 250:
		fmt.Printf("%d is less than 250 \n", n)
	}

}
