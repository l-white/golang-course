//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var color string = "blue"
	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	birthYear, age := 1976, 46
	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial string = "L"
		lastInitial string = "W"
	)
	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	var ageInDays int
	ageInDays = age * 365
	
	fmt.Println("My favorite color is", color + ".")
	fmt.Println("My age is", age, "and my birthyear is", birthYear)
	fmt.Println("My first initial is", firstInitial, "and my last initial is", lastInitial)
	fmt.Println("My age in days is", ageInDays)
}
