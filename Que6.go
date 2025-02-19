// Implement a Go program to convert a given decimal number to binary using loops and conditionals.

package main
import "fmt"

func main(){
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num < 0 {
        fmt.Println("Please enter a non negative number.")
        return
    }

    binary := ""

    if num == 0{
        binary = "0"
    } else{
        for num > 0{
            rem := num % 2
            if rem == 1{
                binary =  "1" + binary
            } else{
                binary = "0" + binary 
            }
            num = num / 2
        }     
    }

    fmt.Println("Binary representation:", binary)
}


/*
package main

import (
	"fmt"
)

func main() {
	var num float64
	fmt.Print("Enter a decimal number: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Please enter a non-negative number.")
		return
	}

	// Extract integer and fractional parts
	intPart := int(num)
	fracPart := num - float64(intPart)

	// Convert integer part to binary
	binaryInt := ""
	if intPart == 0 {
		binaryInt = "0"
	} else {
		for intPart > 0 {
			rem := intPart % 2
			if rem == 1 {
				binaryInt = "1" + binaryInt
			} else {
				binaryInt = "0" + binaryInt
			}
			intPart = intPart / 2
		}
	}

	// Convert fractional part to binary
	binaryFrac := ""
	precision := 5 // Number of decimal places in binary representation

	if fracPart > 0 {
		binaryFrac = "."
		for i := 0; i < precision; i++ {
			fracPart *= 2
			intBit := int(fracPart)
			if intBit == 1 {
				binaryFrac += "1"
				fracPart -= float64(intBit)
			} else {
				binaryFrac += "0"
			}

			if fracPart == 0 {
				break
			}
		}
	}

	// Print final binary representation
	fmt.Println("Binary representation:", binaryInt+binaryFrac)
}

*/