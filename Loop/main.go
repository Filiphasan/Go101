package main

import "fmt"

func main() {
	// for loop
	number := 10
	fmt.Println("Enter a number")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= number; i++ {
		fmt.Println(i)
	}

	// multi for loop
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Print(i*j, " ")
		}
	}

	// for loop with break
	fmt.Println("For loop with break 5")
	for i := 1; i <= number; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// for loop with continue
	fmt.Println("For loop with continue 5")
	for i := 1; i <= number; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// while loop
	fmt.Println("While loop")
	i := 1
	for i <= number {
		fmt.Println(i)
		i++
	}
}
