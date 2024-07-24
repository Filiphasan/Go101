package main

import "fmt"

func main() {
	fmt.Println("Enter count of numbers")
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}
	var numbers = make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Println("Enter number")
		_, err := fmt.Scan(&numbers[i])
		if err != nil {
			continue
		}
	}
	fmt.Println("Sum is: ", sum(numbers...))
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
