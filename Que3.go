// Develop a program to determine if a given year is a leap year using control flow.

package main
import "fmt"
func main(){
	var year int
	fmt.Print("Enter a year:")
	fmt.Scanln(&year)

	if year % 4 == 0 {
		if year % 100 == 0 {
			if year % 400 == 0 {
				fmt.Println(year, "is a leap year")
				} else {
					fmt.Println(year, "is not a leap year")
					}
					} else {
						fmt.Println(year, "is a leap year")
						}
						} else {
							fmt.Println(year, "is not a leap year")
							}

}