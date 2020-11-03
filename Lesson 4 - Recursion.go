package main

import (
	"fmt"
	"math"
)

func main() {
	// Redundancy and how to avoid it
	var testRadius float64 = 3.33

	// fmt.Println(sphereVolume(testRadius))

	var volume float64 = sphereVolume(testRadius)
	fmt.Println(volume)

	num := 10

	for num > 0 {
		num--
	}

	fmt.Println(factorial(6))

	concatenate("Yee", "haww")

	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}

	var floatNum float64 = 0.816732819073

	fmt.Println(getSign(floatNum))
}

func sphereVolume(radius float64) float64 {
	volume := 4 / 3.0 * math.Pi * math.Pow(radius, 3)
	return volume
}

func factorial(num int) int {
	if num < 1 {
		return 1
	}

	return num * factorial(num-1)
}

/*
6 * factorial(5)
6 * (5 * factorial(4))
6 * (5 * (4 * factorial(3)))
6 * (5 * (4 * (3 * (factorial(2))))
6 * (5 * (4 * (3 * (2 * factorial(1)))))
6 * (5 * (4 * (3 * (2 * 1))))
6 * (5 * (4 * (3 * 2)))
6 * (5 * (4 * 6))
6 * (5 * 24)
6 * 120
720
*/

// RECURSION - when a function calls on itself making a loop
func decrement(num int) {
	if num > 0 {
		num--
		decrement(num)
	}
}

func concatenate(string1 string, string2 string) {
	concatenatedString := fmt.Sprintf("%s %s", string1, string2)
	fmt.Println(concatenatedString)
}

func getSign(num float64) string {
	if num > 0 {
		return "Num entered is Positive."
	} else if num < 0 {
		return "Num entered is Negative."
	} else {
		return "Num entered is Zero."
	}
}
