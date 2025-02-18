//Implement a number guessing game where the program generates a random number,
//and the user has to guess it with hints (higher/lower).j

package main

import ("fmt"
		"math/rand"
	)

func main(){
	var guess int;
	var max int;

	fmt.Println("I'll hold a number you need to find it.");
	fmt.Println("Set the maximum range")
	fmt.Scan(&max)

	
	num:= rand.Intn(max)

	for {
		fmt.Println("Guess the number bro:")
		fmt.Scan(&guess)

		if guess == num {
			break;
		} else {
			if num < guess {
				fmt.Println(guess," is lesser than my no.")
			} else {
				fmt.Println(guess," is greater than my no.")
			}
		}
	}
}