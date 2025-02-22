package main

import "fmt"

const a = 10

var p = 100

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println("Sum is:", z)
	}
	add(5, 6)
	add(p, a)
}

func main() {
	call()
	fmt.Println(a)
}

func init() {
	fmt.Println("Hello, World!")
}

/*
2 phases:
1.compilation phase
2.execution phase

***compilation phase***
**code segment**
a = 10
call =func() { ..}
add =func() { ..}
main =func() { ..}
init =func() { ..}

go run main.go=>compile it => main =>./main
go build main.go=>compile it=> main
*/

/*binary file creation*/
