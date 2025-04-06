package main

import "fmt"

var arr2 = [3]string{"my", "name", "is"}

func main() {
	// var arr [5]int
	// arr[4] = 6
	// arr[0] = 9
	// arr[1] = 8
	// arr[2] = 7
	// arr[3] = 5
	arr := [5]int{0, 1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr2[0])
}
