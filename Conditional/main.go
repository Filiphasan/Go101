package main

import (
	"fmt"
)

func main() {
	// if statement
	var number int = 1
	fmt.Println("What's your number?")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("This is not a number", err.Error())
		return
	}

	if number > 0 {
		fmt.Println("Your number is positive")
	} else if number < 0 {
		fmt.Println("Your number is negative")
	} else {
		fmt.Println("Your number is zero")
	}

	// switch statement
	switch number {
	case 1:
		fmt.Println("Your number is 1")
	case 2:
		fmt.Println("Your number is 2")
	case 3:
		fmt.Println("Your number is 3")
	default:
		fmt.Println("Your number is not 1, 2, 3")
	}

	// switch with multi case
	switch number {
	case 1, 2, 3:
		fmt.Println("Your number is one of 1, 2, 3")
	default:
		fmt.Println("Your number is not one of 1, 2, 3")
	}

	// switch with condition
	switch {
	case number > 0:
		fmt.Println("Your number is positive")
	case number < 0:
		fmt.Println("Your number is negative")
	default:
		fmt.Println("Your number is zero")
	}
}
