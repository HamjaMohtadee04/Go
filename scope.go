package main

import "fmt"

var a = 1
var b = 2

func add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	var p = 3
	var q = 4
	add(p, q)
	add(a, b)
	add(a, p)
}
