package main

import "fmt"

var a = 10

func main() {
	fmt.Println(a)
}
func init() {
	// fmt.Println("Init function(it will be excuted first before main.first inti function will be excuted then main function)")
	fmt.Println(a)
	a = 20
	// fmt.Println(a)
}
