package main

import "fmt"

func main() {

	number := 10
	pNumber := &number
	fmt.Println("Pointer:", pNumber)
	fmt.Println("Pointer Value:", *pNumber)

	*pNumber = 20
	fmt.Println("Number:", number)

	number2 := 1000000000
	for num := range number2 {
		*pNumber = num
	}
}
