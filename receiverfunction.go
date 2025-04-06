package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func printUserDetail(usr User) {
	fmt.Println("Name", usr.Name)
	fmt.Println("age", usr.Age)
}

//receiver function
func (usr User) PrintDetails() {
	fmt.Println("Name", usr.Name)
	fmt.Println("age", usr.Age)
}

func (usr User) call(a int) {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
	fmt.Println(a)
}

func main() {
	// var User1 User
	User1 := User{
		Name: "nafi",
		Age:  30,
	}

	User2 := User{
		Name: "hamja",
		Age:  25,
	}

	// printUserDetail(User1)
	User1.PrintDetails() //calling by receiver function
	// printUserDetail(User2)

	User2.call(24)

}
