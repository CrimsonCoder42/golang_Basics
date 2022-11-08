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

func nameGreet(name string) {
	fmt.Println("Hello", name, "welcome to golang")
}

func tester() string {
	return "This is a test of the Devin broadcast system."
}

func addThree(a, b, c int) int {
	return a + b + c
}

func anyNum() int {
	return 13
}

func anyTwoNum() (int, int) {
	return 3, 7
}

func main() {

	nameGreet("Devin")
	fmt.Println(tester())
	var numOne, numTwo = anyTwoNum()
	fmt.Println(addThree(anyNum(), numOne, numTwo))

}
