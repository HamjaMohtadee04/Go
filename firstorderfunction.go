/*
first order function
1.standard function or named function
2.anonymous function
3.immediate function Expression
4.FUNCTION EXPRESSION
*/

/*
1.object(people,animal)
2.property(name,age,COLOR,STUDENT)
3.RELATION()

nafi is a student
apple is a red
sakib is taller than nafi


//first order logic
rule: all customers must pay their pizza bills.
		all studednts must wear their uniform (first order logic works on object,propertu a )

		//higher order logic
		any rule that applies to all customers must also apply to student/nafi

		a rule: all customers must pay tips to the waiters

*/
/*
any of the following 3
1.parameter -> function
2. FUNCTION RETURN
3. BOTH
*/
package main

import "fmt"

//op func(p int, q int) CALLED CALLBACK FUNCTION
// func processOperation(a int, b int, op func(p int, q int)) {
// 	op(a, b)
// }

func call() func(x int, y int) {
	return add
}

func add(x int, y int) {
	sum := x + y
	fmt.Println("Sum is:", sum)

}

func main() {
	// processOperation(5, 6, add)
	sum := call() //function expression
	sum(5, 6)

	// a:= func(){} (also i can store function in a variable)
	// a := true, 45,45.5 ( //here i can store multiple data type this called first class citizen)
	// fmt.Println(a)
}
