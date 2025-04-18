package main

import "fmt"

func changSlice(p []int) []int {
	p[0] = 10
	p = append(p, 11)
	return p
}

//variadic function
func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

// func variadic() {
// 	print(5, 6, 7, 8, 9, 10) //ami amr issa moto argument pathate parbo.
// }
func main() {

	// arr := [6]string{"this", "is", "a", "go", "interview", "question"}
	// fmt.Println(arr)

	// //slice
	// s0 := arr[1:4] // "is" "a" "go"
	// fmt.Println(s0)

	// s1 := s0[1:2] //pointer=1,length=1[1:2 means, 1 theke 1],capacity=4 ["a","go","interview","question"],total=4]
	// fmt.Println(s1)
	// fmt.Println(len(s1))
	// fmt.Println(cap(s1))
	// s := []int{1, 2, 3} //slice literal
	// fmt.Println(s, len(s), cap(s))

	// //slice with built in make() function
	// z := make([]int, 3)
	// z[0] = 5
	// fmt.Println(z, len(z), cap(z))

	// //slice with len and capacity
	// x := make([]int, 3, 5)
	// x[0] = 5
	// x[2] = 4
	// fmt.Println(x, len(x), cap(x))

	// //empty or nil slice
	// var g []int //[],len=0,cap=0
	// fmt.Println(g)

	// //append element in empty slice
	// // g = append(g, 1) //[],len=1,cap=1
	// g = append(g, 1, 2, 3) //[],len=3,cap=3
	// fmt.Println(g)

	// var x []int      //[],len=0,cap=0
	// x = append(x, 1) //[1],len=1,cap=1
	// x = append(x, 2) //[1,2],len=2,cap=2
	// x = append(x, 3) //[1,2,3],len=3,cap=4 {slice len same hole 2 diye gun hoi..jar jonno 123 append hobe kintu cap hobe 4}
	// // x = append(x, 4) //[1,2,3,4],len=4,cap=4 jehetu len ar cap same na so sehetu new arr declare hobe na , karon age ekta cell khali ase tai okhane push korbe

	// y := x //[1,2,3],len=3,cap=4

	// x = append(x, 4) //[1,2,3,4],len=4,cap=4 jehetu len ar cap same na so sehetu new arr declare hobe na , karon age ekta cell khali ase tai okhane push korbe

	// y = append(y, 5) //[1,2,3,5],len=4,cap=4

	// x[0] = 10 // [10,2,3,5]

	// fmt.Println(x) // [10,2,3,5]
	// fmt.Println(y) // [10,2,3,5]

	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := changSlice(a)

	fmt.Println(x)      //[1,2,3,4,10,6,7]
	fmt.Println(y)      //[10,6,7,11]
	fmt.Println(x[0:8]) //[1,2,3,4,10,6,7,11] ektu jor kora hoye gese.jehetu length 10 sehetu 11 print korte prob hobe na.
	print(5, 6, 7, 8, 9, 10)

}

/*
1.compilation phase(compile time)
2.execution phase(run time)


**code segment**
main = func () {...}
*/
/*
1.slice from an array
2.slice from an slice
3.slice literal
4.make func with len
5.make func with len and capacity
6.empty or nil slice
7.slice underlying array rule => 1024 -> increase 100% ,if greater than 1024 then increase 25%
*/
