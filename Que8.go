// Write a program to generate the first N Fibonacci numbers using a loop.

package main
import "fmt"
 func main(){
	var num int

	fmt.Print("Enter the number of terms: ")
	fmt.Scanln(&num)

	if num <= 0 {
        fmt.Println("Please enter a positive integer.")
        return
    }

	t1 := 0
	t2 := 1

	fmt.Println("Fibonacci Series is: ")

	for i:=0;i<num;i++{
		fmt.Print(t1 , " ")

		next := t1 + t2
		t1 = t2
		t2 = next

	}
	fmt.Println()
 }