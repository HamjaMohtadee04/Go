package main
import "fmt"

func main(){
	// age:= 15

	// if age > 15 {
	// 	fmt.Println("You can ride alone")
	// }else if age <=15
	//  18 {
	// 	fmt.Println("You can ride with a guardian")		
	// }else if age ==10{
	// 	fmt.Println("You can't ride")
	// }
	// age := 18
	// sex:= "male"
	// if age ==20 && sex == "male"{
	// 	fmt.Println("You can ride alone")
	// }

	// if age >30 || sex == "male"{
	// 	fmt.Println("You can ride alone")
	// }
	// isPretty := false
	// if !isPretty{
	// 	fmt.Println("You are not pretty")
	// }
	a:=3
	switch a{	
	case 1:
		fmt.Println("One")
	case 2,3:
		fmt.Println("Two or Three")
	default:
		fmt.Println("a is neither 1 nor 2")	
	}
}