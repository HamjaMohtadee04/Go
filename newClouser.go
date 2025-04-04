package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("age =", age)

	//return 5
	show := func() {
		money = money + a + p
		fmt.Println("money", money)
	}
	return show
}

func call() {
	incr1 := outer()
	fmt.Println(incr1) //first money=100+10+100=210
	// incr1()
	// incr1()

	// incr2 := outer()
	// // incr2()
	// incr2()
	// incr2()
	// incr1()
}

func main() {
	call()
}

func init() {
	fmt.Println("bank")
}
