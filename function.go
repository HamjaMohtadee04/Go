package main
import "fmt"
func add(a int, b int) int{
	 sum  := a+b
	return sum
}
func getNumbers(a int , b int) (int,int){
	sum := a+b
	mul := a*b
	return sum,mul
}
 func main(){
 a:=10
 b:=20
// sum := add(a,b)
// fmt.Println(sum) 
 p,q := getNumbers(a,b)
 fmt.Println(p,q)
}