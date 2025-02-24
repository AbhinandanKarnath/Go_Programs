/*   Implement a grading system that takes marks as input and assigns grades based on  
	 predefined categories using conditionals. */

package main
import "fmt"
func main(){
	var  n int
	var avg float64
	tot := 0
	fmt.Print("Enter the number of subjects: ")
	fmt.Scan(&n)
	marks := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter marks for each subject %d :", i+1)
		fmt.Scan(&marks[i])
	}

	for i := 0; i < n ; i++{
		tot = tot + marks[i]
	}

	fmt.Printf("Total marks are: %d\n", tot)

	avg = float64(tot)/float64(n)

	fmt.Printf("Average marks are: %.2f\n",avg)

	if avg >= 90{
		fmt.Print("Passed the examination with Grade A\n")
	} else if avg >= 80 && avg <= 89{
		fmt.Print("Passed the examination with Grade B\n")
	} else if avg >= 70 && avg <= 79{
		fmt.Print("Passed the examination with Grade C\n")
	} else if avg >= 60 && avg <= 69{
		fmt.Print("Passed the examination with Grade D\n")
	} else if avg >= 50 && avg <= 59{
		fmt.Print("Passed the examination with Grade E\n")
	} else if avg >= 40 && avg <= 49{
		fmt.Print("Passed the examination with Grade F\n")
	} else{
		fmt.Print("Sorry, You are failed in the examination!")
	}

}