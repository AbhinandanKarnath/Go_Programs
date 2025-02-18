// Write a Go program to find the sum of digits of a number using a loop.

package main
import (
		"fmt" 
		"math"
)
func main(){
	var num,d int
	sum := 0
	fmt.Print("Enter a number:")
	fmt.Scanln(&num)
	num = int(math.Abs(float64(num)))
	
	for num>0{
		d = num%10
		sum = sum + d
		num = num/10
	}
	fmt.Printf("Sum is %d", sum)
}