package main

import "fmt"

func main() {
	var a, b int
	var op string
for{
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	fmt.Println("Enter the operation to be performed:")
	fmt.Println("1. Addition\n2. Difference\n3. Multiplication\n4. Division\n5. Exit")
	fmt.Scan(&op)

	switch op {
	case "1":
		sum := a + b
		fmt.Printf("Sum of %d and %d is %d.\n", a, b, sum)
		break

	case "2":
		diff := a - b
		fmt.Printf("Difference of %d and %d is %d.\n", a, b, diff)
		break

	case "3":
		prod := a * b
		fmt.Printf("Product of %d and %d is %d.\n", a, b, prod)
		break

	case "4":
		if b == 0 {
			fmt.Println("Error! Division by zero is not allowed.")
		} else {
			div := float64(a) / float64(b)
			fmt.Printf("Division of %d by %d is %.2f\n", a, b, div)
		}
		break

	case "5":
		fmt.Println("Exiting the program.")
		return

	default:
		fmt.Println("Invalid option! Exiting.")
	}
}
}
