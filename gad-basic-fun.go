package main

import (
	"fmt"
)

func addNumbers(number1 int, number2 int) {
	var sum = number1 + number2
	fmt.Println(sum)
}

/*
If consecutive parameters are of the same type, we can avoid writing
the type each time and it is enough to be written once at the end.ie
price int, no int can be written as price, no int.
Example: func calculateBill(price, productNo int) int {
*/
func calculateBill(price int, productNo int) int {
	var totalPrice = price * productNo
	return totalPrice
}

// Multiple return values:
func rectangleProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length * width) * 2
	return area, perimeter
}

func main() {
	fmt.Println("********************************************************")

	// 1st Usage:
	addNumbers(34, 66)
	addNumbers(4222, 78)

	// 2nd Usage: (A)
	fmt.Println("The calculated value:  ", calculateBill(400, 10))

	// 2nd Usage: (B)
	thePrice, prodNo := 50, 10
	theTotalPrice := calculateBill(thePrice, prodNo)
	fmt.Println("The new price is :  ", theTotalPrice)

	fmt.Println("********************************************************")

}
