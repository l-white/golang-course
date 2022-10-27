//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greet(name string) string {
	return "hello, " + name
}

func msg() string {
	return "Greetings!"
}

func addThree (num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func returnNum() int {
	return 7
}

func returnTwoNumbers() (int, int) {
	return 1, 2
}

func returnFuncs() int {
	return addThree(1, 2, 3) + returnNum()
}

func main() {
	fmt.Println("Hello world!")
	fmt.Println(greet("Laura"))
	fmt.Println(msg())
	fmt.Println(addThree(1, 2, 3))
	fmt.Println(returnNum())
	fmt.Println(returnTwoNumbers())

	fmt.Println(returnFuncs())
}
