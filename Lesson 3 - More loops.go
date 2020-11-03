package main

import (
	"fmt"
	"math"
)

func main() {

	// Sum of first n natural numbers
	// 1 + 2 + 3 + 4 + 5

	/*
		var n int = 5
		var sum int = 0
		for i := 1; i <= n; i++ {
			sum += i
		}s

		fmt.Println(sum)

		// Proof of sum of first n natural numbers: (n * (n + 1)) / 2
		// Concatenation Sprintf or Printf
		concatenated := fmt.Sprintf("%s %t", "Function working =", sum == (n*(n+1))/2)
		fmt.Println(concatenated)

		// LOOP REVIEW
	*/

	// %t = boolean, and %d = number value (int), and %f = decimal value (float64)

	/*
		concatenated := fmt.Sprintf("%s %f", "Value =", 1/2.0)
		fmt.Println(concatenated)

		// Scope (where a variable is defined and can be used)
		for i := 1; i < 10; i++ {
			fmt.Println(i)
		}

		// can't print i right here (out of loop's scopes)
	*/

	// Subroutine, is a code block designed to accomplish a specific task
	// It does not return anything but does a task

	subRoutine("Print this - Hello World", 3)

	// Functions, is a code block designed to accomplish a specific task
	// It returns information to where it was called from

	num := 10.35
	squaredNum := squareNum(num)

	fmt.Println(squaredNum)
	resultOfSquaredNum := fmt.Sprintf("%s %f", "The squared num of 10.35 =", squaredNum)

	fmt.Println(resultOfSquaredNum)
}

// This subroutine takes in a placeholder "text" which is a STRING type
func subRoutine(text string, num int) {
	fmt.Println(text)
	fmt.Println(num)
}

// the function squareNum is a function that returns a squared value of a number
func squareNum(num float64) float64 {
	// num *= num
	num = math.Pow(num, 2)

	return num
}
