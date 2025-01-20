package main

import "fmt"

var a = 1
var b = 2

func main() {
	x := 3

	if x >= 2 {
		p := 4
		fmt.Println("x is greater than or equal to 2")
		fmt.Println("i have", p, "goats")
	}
	// fmt.Println("i have", p, "goats") -> p is not define locally or globally so it will throw an error
}
