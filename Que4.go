// Write a Go program to calculate and print the sum of all even numbers from 1 to N using a loop.

package main
import "fmt"
func main(){
	var n int
	sum := 0
	fmt.Print("Enter a number:")
	fmt.Scanln(&n)
	for i:=0; i<=n; i++{
		if(i%2==0){
			sum += i
		}
		
	}
	fmt.Println("Sum of even numbers:", sum)
}