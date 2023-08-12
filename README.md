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

