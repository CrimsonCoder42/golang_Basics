//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 51; i++ {
		sum++
		if sum%3 == 0 && sum%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if sum%5 == 0 {
			fmt.Println("Buzz")
		} else if sum%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
