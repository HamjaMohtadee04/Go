package main

import "fmt"

var a = 10

func main() {
	age := 30
	if age > 18 {
		a := 20 // variable shadowing
		fmt.Println(a)
	}
	fmt.Println(a)
}
