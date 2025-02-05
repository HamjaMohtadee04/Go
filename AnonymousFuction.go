package main

import "fmt"

func add(a int, b int) { //standard or name function
	sum := a + b
	fmt.Println("Sum is:", sum)

}

func main() {
	// what is Expression
	a := 10 //full line is an expression

	//if expression
	if a > 0 { //if block
		fmt.Println("A is greater than 0")
	}

	add(5, 6) //calling standard function ,called means invoked the function

	func(a int, b int) { //anonymous function which will not directly stored in go and also Immediately Invoked Function Expression or IIFE, it should be called immeidately after declaration called iife or anonymous function
		c := a + b
		fmt.Println("Sum is:", c)
	}(10, 20) //immediately called after declaration
}

func init() {
	fmt.Println("Init function")
}
