package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Breaking the outer for loop
	// A go program to generate a random even number

randomLoop:
	for {
		switch k := rand.Intn(200); {
		case k%2 == 0:
			fmt.Printf("The generated even number is: %d", k)
			break randomLoop
		}
	}
}
