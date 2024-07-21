package main

import "fmt"

func main() {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}

	for i := 1; i <= number; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
