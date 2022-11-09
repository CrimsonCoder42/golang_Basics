//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice(sides int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(sides) + 1
	return randNum
}

func multiRolls(roll, sides int) int {
	count := 0
	for i := 0; i < roll; i++ {
		randNum := rollDice(sides)
		count += randNum
	}
	return count
}

func snakeEyes(roll, dice int) bool {
	if roll == 2 && dice == 2 {
		return true
	}
	return false
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func main() {

	numOfSides := 6
	numOfDice := 2
	numOfRolls := 2
	totalRolls := numOfDice * numOfRolls

	rollTotal := multiRolls(totalRolls, numOfSides)
	//* Print the sum of the dice roll
	if snakeEyes(rollTotal, numOfDice) {
		fmt.Println("Snake eyes")
	} else if rollTotal == 7 {
		fmt.Println("Lucky 7", rollTotal)
	} else if isEven(rollTotal) {
		fmt.Println("Even", rollTotal)
	} else if !isEven(rollTotal) {
		fmt.Println("Odd", rollTotal)
	} else {
		fmt.Println(rollTotal)
	}
}
