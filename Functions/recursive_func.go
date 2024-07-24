package main

import "fmt"

func main() {
	fmt.Println("Enter number")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}
	fmt.Println("The factorial is: ", factorial(number))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// anon
