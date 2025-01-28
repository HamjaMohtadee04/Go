package main

import "fmt"

func add(a int, b int) { //jei function er nam ase tai named function/standard function
	res := a + b
	fmt.Println(res)
}

func main() {
	add(10, 20)
}
