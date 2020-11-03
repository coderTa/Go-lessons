package main

// Budgeter

/* Asks for your monthly income and your daily or monthly expenses
and outputs how much money you earned and spent, and whether you
made any savings or if you are now in debt. */

import (
	"fmt"
	"math"
	"math/rand"
)

var monthLength float64 = 29.0 - float64(rand.Intn(2))

/* Tell user what the program does with intro() method
and call methods that prompt them to provide input and then output results */

func main() {
	intro()

	var totalIncome float64 = returnIncome()
	var totalExpenses float64 = returnExpense()

	fmt.Println(totalIncome)
	fmt.Println(totalExpenses)
	finisher(totalIncome, totalExpenses)
}

// intro(), should ouput what the budgeter does

func intro() {
	fmt.Println("Welcome to the Budgeter. Be ready to insert your montly income and expenses.")
	fmt.Println("It will also tell you your net montly income.")
}

/* returnIncome(), asks for how many sources of montly income and adds the amount
of money in each category before returning it as a float64. */

func returnIncome() float64 {
	var totalIncome float64
	var incomeCategories int

	fmt.Print("How many categories of income? ")
	fmt.Scanln(&incomeCategories)

	for i := 0; i < incomeCategories; i++ {
		fmt.Print("\tNext income amount? $")
		var incomeAmount float64
		fmt.Scanln(&incomeAmount)
		totalIncome += incomeAmount
	}

	fmt.Println()

	return totalIncome
}

/* returnExpenses(), asks whether the sources of expenses you're about to provide are
montly or daily and then asks for how many sources of expenses and adds the amount
of money in each category. It converts expenses to monthly if daily option was chosen
before returning expenses as a float64. */

func returnExpense() float64 {
	fmt.Print("Enter 1) monthly or 2) daily expenses? ")

	var totalExpenses float64
	var expenseType int

	fmt.Scanln(&expenseType)

	fmt.Print("How many categories of expenses? ")

	var expenseCategories int
	fmt.Scanln(&expenseCategories)

	for i := 0; i < expenseCategories; i++ {
		fmt.Print("\tNext expense amount? $")
		var expenseAmount float64
		fmt.Scanln(&expenseAmount)
		totalExpenses += expenseAmount
	}

	if expenseType == 2 {
		totalExpenses *= monthLength
	}

	fmt.Println()

	return totalExpenses
}

/* finisher(), outputs how much you earned and spent in the month and

 */

func finisher(totalIncome float64, totalExpenses float64) {
	var roundedIncome float64 = roundNum(totalIncome, 2)
	var roundedExpenses float64 = roundNum(totalExpenses, 2)
	var leftoverFunds float64 = roundNum(roundedIncome-roundedExpenses, 2)

	fmt.Println(fmt.Sprintf("Total income of $%f ($%f/day)", roundedIncome, roundNum(roundedIncome/monthLength, 2)))
	fmt.Println(fmt.Sprintf("Total expenses of $%f ($%f/day)", roundedExpenses, roundNum(roundedExpenses/monthLength, 2)))
	fmt.Println()
	fmt.Printf("You have $%f leftover funds this month.", leftoverFunds)
	fmt.Println()

	if leftoverFunds <= 0 {
		fmt.Println("RIP you're broke lmao")
	}
}

/* roundNum(), takes a number and number of digits to round to and rounds the provided number
before returning the float. */

func roundNum(num float64, placeVal float64) float64 {
	var roundedNum float64
	roundedNum = math.Floor(num*math.Pow(10, placeVal)+0.5) / math.Pow(10, placeVal)

	return roundedNum
}
