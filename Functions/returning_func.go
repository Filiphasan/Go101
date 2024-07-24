package main

import "fmt"

func main() {
	var firsNumber int
	var secondNumber int

	fmt.Println("Enter the first number")
	_, err := fmt.Scan(&firsNumber)
	if err != nil {
		return
	}
	fmt.Println("Enter the second number")
	_, err2 := fmt.Scan(&secondNumber)
	if err2 != nil {
		return
	}

	result := addition(firsNumber, secondNumber)
	add, subt := additiondAndSubtract(firsNumber, secondNumber)
	mult, div := multiplyAndDivide(firsNumber, secondNumber)
	fmt.Println("Result is: ", result)
	fmt.Println("Addition is: ", add)
	fmt.Println("Subtraction is: ", subt)
	fmt.Println("Multiplication is: ", mult)
	fmt.Println("Division is: ", div)
}

// single return
func addition(first int, second int) int {
	return first + second
}

// multiple return
func additiondAndSubtract(first int, second int) (int, int) {
	return first + second, first - second
}

// Named return
func multiplyAndDivide(first int, second int) (mult int, div int) {
	mult = first * second
	div = first / second
	return
}
