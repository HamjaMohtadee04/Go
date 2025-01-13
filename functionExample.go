package main

import "fmt"

func printWelcomeMessage() {
	fmt.Println("Welcome to the function example")
}

func getUserName() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func getNumbers() (int, int) {
	var a int
	var b int
	fmt.Println("enter first number: ")
	fmt.Scanln(&a)
	fmt.Println("enter second number: ")
	fmt.Scanln(&b)
	return a, b
}
func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}
func display(name string, sum int) {
	fmt.Println("Hello, ", name)
	fmt.Println("Sum = ", sum)
}
func printGoodbyeMessage(name string) {
	fmt.Println("Goodbye, ", name)
	fmt.Println("thank you")
}
func main() {
	// fmt.Println("hello")
	// var name string
	// fmt.Println("Enter your name: ")
	// fmt.Scanln(&name)

	// var a int
	// var b int
	// fmt.Println("enter first number: ")
	// fmt.Scanln(&a)
	// fmt.Println("enter second number: ")
	// fmt.Scanln(&b)
	// sum := a + b

	// // display
	// fmt.Println("Hello, ", name)
	// fmt.Println("Sum = ", sum)

	// //  print a good bye message
	// fmt.Println("Goodbye, ", name)
	// fmt.Println("thank you")
	printWelcomeMessage()
	name := getUserName()
	num1, num2 := getNumbers()
	sum := add(num1, num2)
	display(name, sum)
	printGoodbyeMessage(name)
}
