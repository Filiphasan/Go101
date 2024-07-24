package main

import "fmt"

func main() {

	firstNum := func() int {
		return 10
	}
	var secondNum int
	fmt.Println("Enter number")
	_, err := fmt.Scan(&secondNum)
	if err != nil {
		return
	}
	fmt.Println("Sum is: ", additionNumbers(firstNum(), secondNum))
}

func additionNumbers(first int, second int) int {
	return first + second
}
