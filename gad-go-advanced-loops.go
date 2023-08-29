package main

import "fmt"

func main() {
	fmt.Println("Hold on!")
	for i := 0; i < 3; i++ {
		fmt.Printf("Numbers %d \n", i)
	}

	/*
		Use of BREAK statement
		It is used to terminate the for loop abruptly b4 it finishes its normal
		execution and move the control to the line of code just after the for loop.
	*/

	for b := 1; b <= 8; b++ {
		if b > 4 {
			// Here, Loop is terminated if b > 2
			break
		}
		fmt.Printf("%d ", b)
	}
	fmt.Printf("\n Break line after looping.\n")

	/*
		The continue statement:
		It is used to skip the current iteration of the for loop. All code
		present in a for loop after the continue statement will not be executed
		for the current iteration. The loop will move on to the next iteration.
	*/

	for k := 1; k <= 20; k++ {
		if k%2 == 0 { // U can use if k%2 != 0 for even numbers
			continue
		}
		fmt.Printf("%d ", k)
	}

	/*
		There, if i%2 == 0 checks if the remainder of dividing i by 2 is 0.
		If it is zero, then the number is even and continue statement is executed
		and the control moves to the next iteration of the loop. Hence the print
		statement after the "continue" will not be called and the loop proceeds
		to the next iteration. The output: 1 3 5 7 9 11 13 15 17 19

		Nested for loops:
		=================
		Nested loop is a for loop which has another for loop inside it.
	*/

	fmt.Print("\n")
	var n = 5
	for m := 0; m < n; m++ {
		for p := 0; p <= m; p++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	/*
		Labels:
		-------
		Labels can be used to break the outer loop from inside the inner for loop.
	*/

	for g := 0; g < 4; g++ {
		for h := 1; h < 5; h++ {
			fmt.Printf("g = %d , h = %d \n", g, h)
		}
	}
	fmt.Printf("**************************************\n")

	/*
		Let's rewrite the program above using labels
		--------------------------------------------
		What if we want to stop printing when i and j are equal. To do this we
		need to break from the outer for loop. Adding a break in the inner for
		loop when i and j are equal will only break from the inner for loop.
	*/

outer:
	for q := 0; q < 4; q++ {
		for t := 1; t < 5; t++ {
			fmt.Printf("q = %d , t = %d \n", q, t)
			if t == q {
				break outer
			}
		}
	}

}
