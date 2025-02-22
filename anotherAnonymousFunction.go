package main

import "fmt"

var sum2 = func(a int, b int) {
	c := a + b
	fmt.Println("Sum is:", c)
}

func add(a int, b int) { //standard or name function
	sum := a + b
	fmt.Println("Sum is:", sum)

}

func main() {

	sum2(10, 20)
	add(5, 6)

	// a := 10 // variable expression

	//function expression or assign function in variable
	add := func(a int, b int) { //function expression
		c := a + b
		fmt.Println("Sum is:", c)
	}
	add(10, 20) //calling function expression
}

func init() {
	fmt.Println("Init function")
}
