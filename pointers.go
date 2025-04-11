package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	obj1 := User{ //instance
		Name:   "nafi",
		Age:    24,
		Salary: 0.10,
	}

	//printing using pointer
	p := &obj1
	fmt.Println(*p)
	fmt.Println(p.Salary)

	//print using struct
	// fmt.Println(obj1)
	// fmt.Println(obj1.Salary)

}

// func print(numbers *[3]string) {
// 	fmt.Println(numbers)
// }
// func main() {
// 	//pointer means address of memory
// 	// x := 20
// 	// fmt.Println("x", x)
// 	// p := &x //address of

// 	// *p = 30
// 	// fmt.Println("x", x)

// 	// fmt.Println("address:", p)
// 	// fmt.Println("value at the address", *p)
// 	// array := [3]int{4, 5, 6}
// 	// print(&array) //passing the address of the array
// 	array := [3]string{"babe", "miss", "you"}
// 	print(&array)
// }
