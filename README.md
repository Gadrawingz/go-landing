# Go Gad!

## Go basics
Go is an open source programming language designed for building simple, fast, and reliable software.

Go has an easy-to-use and understand syntax while providing top-tier performance.

## Go applications
You can use Go to build various applications such as: 
1. Web apps 
2. Command-line apps
3. Cloud services,
4. Networking.

## Go features
- Generics
- Goroutines
- Concurrency,
- Type inference,
- Garbage collection, etc.

## Local Setup
The following programs have to be set up before getting hands dirty with code:
- **Go** Go distribution actually comes with Go CLI tool.
- **Install an IDE such as VS Code** for writing code. We use VS Code, but  you can use any other code editor of your choice.

# If Visual Studio is used as text editor.
## I. VS Code extensions
> VS Code extensions let you add language support, debugger and tools for easier development.
 #### Go extensions:
> VS Code provides features like code completion, code navigation, etc.
1. Go (Go team at Google)
2. Extension

## If GoLand is used as text editor!
You insall GoLand from jetbrains as official site and follow the given instructions.

## II. Running first go project
>1. Create the your go file:
**main.go**
> 2. Initialize our project module
> go mod init **[module-path]** (creates a new module)
Module path can correspond a repo you plan to publish your module to.
(eg. github.com/Gadrawingz/go-landing)
> *go mod init go-landing*
> ### The nitialized a go.mod file:
> Describes the modules: with **name/module** path and  **go version**  used in program.
> The module path is also the **import path** e.g. import "github.com/gadira/go-landing"

## Project Structure
- **package main**: Every go file must start with the **package name** statement. 
- **import "fmt"**: This package is imported and it will be used inside the main function to print text to the standard output.
- **fmt.Println("Hello World")** : The Println function of the fmt package is used to write text to the standard output.
- **func main()** - main is the special function. The program execution starts from the main function. It should always reside in the main package.
- The **{** and **}** indicate the start and end of the main function.

## Packages in Go
> In go everything is organized into packages.
> We are gonna be using go packages throughout our application
> A **package** is a collection of source code.
> 
> **Packages** are used to provide code compartmentalisation and reusability. Here the package name used is main
> 
> **Note1:** All our code must belong to a package
> **Note2:** The first statement in Go file must be "package...."
> **Go's standard library**, provides different core packages for us to use.
> fmt is the one of these, Which you can use by importing it.

## Where to start the program?
## Where's the entry point?

> An entry point of GO is main function...
> A go program can only have 1 main function, because you can only have one (1) entry point.


# GO DRAFT

## Constants
Constants are like variables, except that their values cannot be changed.

## Print formatted data
(Whebever we are printed our text mixed with vairables: )
> fmt.printf("Some text with a variable %s", myVariable)


# Data types
> In any programming languages you have multiple data types
> Difference: which data types do they support exactly?
>  Strings, Booleans, Arrays, Integers, Maps, etc....

# Different data types
> Each data types can do different things and behaves differently
> Note: Go is a statically typed language: you need to tell Go compiler, the data type when declaring the variables

> In Go, you detect/discover mistakes at compile time, NOT at runtime...

Strings
Integers 
Booleans 
Maps
Arrays

## %T Lets you print any type of variable
---

## Floating Point Types
Numbers that can contain a decimal component
> Use Cases: 
Statistical Data
Monetary Data

## "fmt" package
> Different functions for: 
Formatted input and Output (I/O)
Print messages
Collect user Input
Write into a file

## Pointer (What is a pointer?)
> A pointer is avariable that points to the memory address of another variable.
> In go lang, a pointer is called a special variable.

## Why Using Pointer?


## Scan 
> Scan(userName)
Scan("")  -> Passing the variable

> Scan(&userName)
Scan(0xc00124) -> Passing the reference


#Arrays and Slices in Go

"Peter", "Dan", "Paula"
We want to store the entered data in some kind of a list!

We want to store the entered data in some kind of a list to keep track of 
**Who is attending the conference **  **Who booked the tickets? **  
 > Data structure to store collection or elements in a single variable is:
"Gad", "Dan", "Peter", "Dan", "Paula"
> We want to store the entered user data in some kind of a list.


## Lets create an array for all the booking!
Arrays in go has fixed size
Fixed size = How many elements the array can hold
> var variable_name [size] variable type

# SLICES
Problem with array
bookings = ["Tom", ...........................]
booking 1 user books all the tickets..
having a array with size 50, with only 1 element inside is not sufficient/
#What if we don't know the size when creating it? 

# What is slices in Go?
Slice is an abstraction of an Array
Slices are more flexible and powerful: variable-length or get a sub-array of its own.
Slices are also index-based and have a size, But is resized when needed.

> To define an array, We basically create an array without a size definition 

## Append:
>It adds the element(s) at the end of the slice.
> Grows the slice if a greater capacity is needed and returns the updated value.
> Adding value to a slice is different than array but retrieving the value is the same..

# Loops in Go

## For each loop
bookings = ["Bob Marley", "Micheal Jack"]
> For each element in the list, we want to executre the same logic!
We want to have this, To print only the first names.
firstNames = ["Bob", "Micheal"]


> **Range** iterate over elements for different elements for different data structures. So, Not only arrays and slices.

> For arrays and slices , range provides the index and value for each element.

> "strings.Fields()"
Splits the string with white space as separator.
"Nicole Smith"    and ["Nicole", "Smith"]



