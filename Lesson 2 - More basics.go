package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	var num int = 42
	// greeting := "Hello"

	/* Conditional statements
	if, if-else, else-if (elif)
	==, equality, not = ASSIGN
	!=, not equal
	>, greater than
	<, less than
	&&, and operator, requires that input is true on left and right
	||, or operator, requires that one input is true on left or right
	*/

	if num == 42 {
		fmt.Println("Meaning of Life?")
	}

	if num != 25 {
		fmt.Println("Not meaning of life.")
	}

	if num > 0 {
		fmt.Println("The meaning of life is greater than 0.")
	}

	if num < 100 {
		fmt.Println("The meaning of life is less than 100.")
	}

	// Multiple conditions

	if num < 100 && num > 0 {
		fmt.Println("Did you know that the meaning of life is between 0 and 100?")
	}

	if num == 3 || num == 69 {
		fmt.Println()
	}

	// Cases, different things you want to happen when different conditions are met
	// if-else statement, two cases, the second case does not have a condition
	// else-if statement, as many cases as you want, checks cases in order, stops checking when a condition is met

	if num > 50 {
		fmt.Println("Num is greater than 50")
	} else if num < 0 {
		fmt.Println("Num is smaller than 0")
	} else {
		fmt.Println("Num is not greater than 50 or smaller than 0")
	}

	// LOOPS

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// While loops (normally called "while loop" in other languages)

	var tally int

	for num < 1000 {
		num += 10
		tally++

		fmt.Println(tally)
		fmt.Println(num)
	}

	var prev int = 0

	// FIBONACCI??? NOPE

	/*for i := 0; i <= 100; i += prev {
		fmt.Println(i)
		prev += i
	}
	*/
}
