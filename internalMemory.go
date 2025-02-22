package main

import "fmt"

var (
	a = 10
)

func add(x int, y int) {
	sum := x + y
	fmt.Println("Sum is:", sum)

}

func main() {
	add(5, 6)
	add(a, 6)
}

func init() {
	fmt.Println("Hello, World!")
}
