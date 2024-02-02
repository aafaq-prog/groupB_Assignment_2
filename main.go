package main

import (
	"fmt"
)

func main() {
	// Factorial
	fmt.Print("1.Fictorial -   \n")
	fmt.Print("author name:Smit Sandeepkumar Shah Student ID:500221322 \n")
	num := 5
	factorialResult := factorial(num)
	fmt.Printf("Factorial of %d: %d\n", num, factorialResult)
	fmt.Println("-------------------------------------------------------------------")

	fmt.Print("8.check whether the value of variable is greater than or less than 5 \n")
	fmt.Print("author name:Hirenkumar Savani   Student ID:500226947 \n")
	// conditional statements to check whether the value of variable is greater than or less than 5
	var number1 int
	var number2 int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number1)
	fmt.Scan(&number2)
	greaterOrLess(number1, number2)
}

// author name:Smit Sandeepkumar Shah
// Student ID:500221322
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func greaterOrLess(number1 int, number2 int) {
	if number1 > number2 {
		fmt.Printf(" %d is greater than %d \n", number1, number2)
	} else {
		fmt.Printf(" %d is less than %d \n", number1, number2)
	}
}
