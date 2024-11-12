package recursion

import "fmt"

// var fact_list []int

func main() {
	var fact int
	fmt.Println("Please enter the number you would like to see the factorial of: ")
	fmt.Scan(&fact)
	fmt.Println(factorial(fact))
	// fmt.Println(fact_list, " = ", fact)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	// fact_list = append(fact_list, number)
	return number * factorial(number-1)


	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// 	fact_list = append(fact_list, i)
	// }
	// return result
}

// factorial of 5 => 5*4*3*2*1 => 120