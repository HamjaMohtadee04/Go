package main

import "fmt"

//code segment1
type User struct {
	Name string //name and age are the member variables of struct or property
	Age  int
}

//code segment2
func main() {
	var user1 User
	user1 = User{ // making of instance can be called instantiate
		Name: "John",
		Age:  30,
	}

	user2 := User{ // instance of struct or object of struct
		Name: "Jane",
		Age:  25,
	}
	fmt.Println("User1:", user1.Name)
	fmt.Println("User1:", user1.Age)
	fmt.Println("User2:", user2.Name)
	fmt.Println("User2:", user2.Age)
}
