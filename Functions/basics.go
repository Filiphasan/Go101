package main

import (
	"fmt"
)

func main() {
	hello()
	helloWithName("Hasan")

	var message string
	fmt.Println("What's your name?")
	_, err := fmt.Scan(&message)
	if err != nil {
		return
	}

	helloWithName(message)
}

func hello() {
	fmt.Println("Hello")
}

func helloWithName(message string) {
	fmt.Println("Hello " + message)
}
