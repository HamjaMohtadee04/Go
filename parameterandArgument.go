package main

import "fmt"

func add(a int, b int) { //parameter =>a,b
	sum := a + b
	fmt.Println("Sum is:", sum)
}

func main() {
	add(5, 6) //argument =>5,6
}
