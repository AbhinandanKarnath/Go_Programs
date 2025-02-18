// Implement a program to check if a number is prime or not using conditionals and loops.

package main
import "fmt"
func main(){
	var n int
	count := 0
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)
    
	for i:=1; i<=n; i++{
		if n%i==0{
			count++
		}
	}
	if count == 2{
		fmt.Println(n, "is a prime number")
	}else{
		fmt.Println(n, "is not a prime number")
	}
}