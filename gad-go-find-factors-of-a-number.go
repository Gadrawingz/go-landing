package main

import "fmt"

func main() {

    var factSum, i int

    fmt.Print("Enter any Number to find Factors = ")
    fmt.Scanln(&factSum)

    fmt.Println("The Factors of the ", factSum, " are = ")
    for i = 1; i <= factSum; i++ {
        if factSum%i == 0 {
            fmt.Println(i)
        }
    }
}
