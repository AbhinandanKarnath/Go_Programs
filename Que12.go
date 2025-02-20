// Create a program to calculate the power of a number using a loop (without using the  math package).

package main
import "fmt"
func main(){
	var num , pow int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	fmt.Print("Enter the power: ")
	fmt.Scanln(&pow)

	result := 1
	for i:=0;i<pow;i++{
		result = result * num
	}
	fmt.Printf("%d to the power %d is %d.",num,pow,result)
}
