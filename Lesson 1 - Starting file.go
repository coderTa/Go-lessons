package main

// importing the package "fmt" & "math"
import "fmt"
import "math"

// function main, meaning that this function is necessary
func main() {
	fmt.Println("Hello, world!")



/* 
This is a
multiline
comment
*/

// Strings = "text"
// Integers = "ints (019301) a whole number"
// Float = "decimal num (3.14) float64"

// Variables can be typed or untyped (can or cannot be changed later) go is typed (cannot be changed later)

// Variable "egg" which is a float (explicit)
var eggs float64 = 10
eggs += 2

// Raising var eggs to a power of 2
fmt.Println(math.Pow(eggs, 2))

// Prints the variable egg and its value
fmt.Println(eggs)

// Another way to define a variable (implicit)
greeting := "Hello"
fmt.Println(greeting)
greeting += ", how are you"
fmt.Println(greeting)

// Assigning two variables at once
cats, dogs := 3, 6
fmt.Println(cats, dogs)

// Defining pi
var pi float64 = math.Pi
fmt.Println(pi)



// PROBLEMS

var sideLength float64 = 4
var area float64 = math.Pow(sideLength, 2)

fmt.Println(sideLength)
fmt.Println(area)

question := "Who you gonna call?"
question += " GHOSTBUSTERS!!!"

fmt.Println(question)

//var radius float64 = 5
//var cicumference float64 = 2 * radius * math.Pi

//text := "The circumference is: "

}