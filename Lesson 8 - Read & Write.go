package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// File tool. Making and writing files easily with the console.
// Whitespace, tab, space, return

func main() {
	var doRun bool = true
	scanner := bufio.NewReader(os.Stdin)

	intro()

	for doRun {
		fmt.Print("(C)reate file, (V)iew file, (D)elete file, (Q)uit? ")

		userInput, _ := scanner.ReadString('\n')
		//c\n
		userInput = strings.TrimSpace(userInput)
		//c\n
		userInput = strings.ToUpper(userInput)
		//C

		// var userInput string
		// fmt.Scanln(&userInput)
		// userInput = strings.ToUpper(userInput)

		if userInput == "C" {
			writeFile(scanner)
		} else if userInput == "V" {
			readFile(scanner)
		} else if userInput == "D" {
			deleteFile(scanner)
		} else if userInput == "Q" {
			doRun = false
		}
	}
}

// intro(), short introduction of the tool to the user

func intro() {
	fmt.Println("Welcome to the file tool.")
	fmt.Println("This tool allows you to create and view files.")
	fmt.Println("What would you like to do?")
	fmt.Println()
}

// readFile(), ask for a file name and output the contents of the file

func readFile(scanner *bufio.Reader) {
	fmt.Print("Please enter the name or path of the file you want to read:")
	fileInfo, _ := scanner.ReadString('\n')
	fileInfo = strings.TrimSpace(fileInfo)

	// data, err := ioutil.ReadFile("\Users\Ben\Documents\Go Tutor\Lesson 8 - Read & Write.go")
	data, err := ioutil.ReadFile(fileInfo)

	if err != nil {
		fmt.Println("File reading error", err)
		readFile(scanner)
		return
	}

	fmt.Println("Contents of file")
	fmt.Println(string(data))
}

/* writeFile(), ask for the name of the file the user wants to create, check if a file with
that name already exists, and if it does, ask if they want to overwrite it, start over
otherwise. After this, asks what the contents of the file should be, and output the size
of the new file in bytes. */

func writeFile(scanner *bufio.Reader) {
	fmt.Print("Please enter the name of the file you want to create: ")
	fileName, _ := scanner.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	if checkFileExistence(fileName) {
		fmt.Println("A file with this name already exists, do you want to overwrite it?")
		var userResponse string = getYesOrNo(scanner)

		if userResponse == "NO" {
			writeFile(scanner)
			return
		}
	}

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		writeFile(scanner)
		return
	}

	fmt.Print("Please enter the contents of the file: ")
	fileContents, _ := scanner.ReadString('\n')
	fileContents = strings.TrimSpace(fileContents)

	length, err := file.WriteString(fileContents)

	if err != nil {
		fmt.Println(err)
		file.Close()
		writeFile(scanner)
		return
	}

	fmt.Println(length, "bytes were written successfully")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		writeFile(scanner)
		return
	}

}

/* checkFileExistence() is a helper function we will use to check if a file with the given name
already exists returning true if it does exist or false if it does not.*/

func checkFileExistence(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}

	return true
}

/* getYesOrNo() is another helper function that will repeatedly ask users for input until they
provide a yes or no answer. The function will return the response of the user in all caps
so that it is not case sensitive. */

func getYesOrNo(scanner *bufio.Reader) string {
	var doRun bool = true
	var userResponse string

	for doRun {
		userResponse, _ := scanner.ReadString('\n')
		userResponse = strings.TrimSpace(userResponse)
		userResponse = strings.ToUpper(userResponse)

		if userResponse == "YES" || userResponse == "NO" {
			doRun = false
		} else {
			fmt.Println("Please enter \"Yes\" or \"No\".")
		}
	}

	return userResponse
}

// deleteFile(), asks for a name and deletes that file

func deleteFile(scanner *bufio.Reader) {
	fmt.Print("Please enter the name or path of the file you want to delete:")
	fileInfo, _ := scanner.ReadString('\n')
	fileInfo = strings.TrimSpace(fileInfo)

	err := os.Remove(fileInfo)

	if err != nil {
		fmt.Println("File reading error", err)
		readFile(scanner)
		return
	}

	fmt.Println("File deleted successfully")
}
