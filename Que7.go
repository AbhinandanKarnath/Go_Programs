// Develop a Go program that calculates the factorial of a number using recursion. 

package main
import "fmt"

func main(){
	var num int
	fmt.Print("Enter the number to calculate factorial: ")
	fmt.Scan(&num)
	fact := factorial(num)
	fmt.Printf("Factorial of %d is %d",num,fact)
}

func factorial( n int)int{

	if n == 0 || n == 1{
		return 1
	}
	return n * factorial(n-1)
	
}