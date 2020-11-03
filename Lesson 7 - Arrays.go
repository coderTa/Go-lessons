package main

import (
	"fmt"
)

// Introduction to arrays and datastructures (0/1 based indexing - starting at 0 or 1)

/*
Default values:

int = 0
float64 = 0.0
bool = false
string = ''
*/

func main() {
	var scores []int
	scores = []int{10, 20, 30}

	// var scores [3]int
	// scores[0] = 11
	// fmt.Println(scores[0])
	fmt.Println(sumOfArray(scores))

	fmt.Println(scores)
	doubleArrayElements(scores)
	fmt.Println(scores)
	reverseArray(scores)
	fmt.Println(scores)

	x := make(map[string]int)
	x["Grade A"] = 42
	x["Grade B"] = 90
	fmt.Println(x)
}

// Sum elements of array and return the int
func sumOfArray(array []int) int {
	var sum int = 0

	// len(array) = number of elements = 5
	// array[0] = 10

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	return sum
}

/* Double each element in order using generic for loop
Arrays use references, there's only ever 1 array for each array that you create
When it is used elsewhere, the variable is simply a pointer. No return needed
because we are directly interacting with the array */

func doubleArrayElements(array []int) {
	// array[0] = 1
	// Numeric for loop
	/*
		for i := 0; i < len(array); i++ {
			array[i] *= 2
		}
	*/

	for index, value := range array {
		array[index] = value * 2
	}
}

// Reverse order to elements in the array using numeric for loop

func reverseArray(array []int) {
	for start, end := 0, len(array)-1; start < end; start, end = start+1, end-1 {
		array[start], array[end] = array[end], array[start]
	}
}

func printMap(x map[string]int) {
	x["Grade A"] = 90
	x["Grade B"] = 80
	for key, value := range x {
		fmt.Print(key + ": ")
		fmt.Println(value)
	}
}
