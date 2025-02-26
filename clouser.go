package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100 //first class order function
	age := 30

	fmt.Println("age =", age)

	// return 5
	show := func() { //outer anonymous function
		money = money + a + p
		fmt.Println("money", money)
	}
	return show //this is a clouser and very important
}

func call() {
	incr1 := outer()
	fmt.Println(incr1)
	// incr1() //its just show function, for first money=100+10+100=210
	// incr1() //for second time money=210+10+100=320
	// incr1() //for third time money=320+10+100=430

	// incr2 := outer()
	// incr2() //for first time money=100+10+100=210,again new memory is allocated and new calculation
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

/*
***code segment***
//compilation phase
a = 10
outer=func() { ..}
outerAnonymous1=func() { ..}
call=func() { ..}
main=func() { ..}
init=func() { ..}
//execution phase
./couser

*/
