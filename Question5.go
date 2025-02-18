// 5. Create a Go program that reads a string input and determines if it is a palindrome.

package main

import "fmt"

func main(){
	var num int
	var str string
	var flag = true

	fmt.Println("Enter a string: ")
	fmt.Scan(&str)

	num = len(str) - 1
	fmt.Println(num)


	for i , j := 0 , num ; i < num/2 && j >= num/2 ; i , j = i+1 , j-1 {
		fmt.Println(i," ",j)
		fmt.Println(str[i]," ",str[j])
		if str[i] != str[j] {
			flag = false
			break;
		}
	}

	if flag {
		fmt.Print(str," is a palindrome")
	} else {
		fmt.Print(str," is not a palindrome")
	}
}