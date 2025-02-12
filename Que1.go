/*Write a Go program that accepts user input for three numbers and determines the
largest using conditional statements. */

package main
import "fmt"
func main() {
	var num1, num2, num3 int
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)
	fmt.Print("Enter the third number: ")
	fmt.Scanln(&num3)
	if num1 >= num2 && num1 >= num3{
		fmt.Println("The largest number is",num1)
	}else if num2 >= num1 && num2 >= num3{
		fmt.Println("The largest number is",num2)
	} else{
		fmt.Println("The largest number is",num3)
	}
	
}