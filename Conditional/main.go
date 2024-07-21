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
}
