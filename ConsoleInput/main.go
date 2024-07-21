package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// bufio package
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What's your name?")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Hello " + name)

	// fmt package Scan function
	fmt.Println("What's your surname?")
	var name2 string
	_, err := fmt.Scan(&name2)
	if err != nil {
		return
	}
	fmt.Println("Hello " + name + " " + name2)

	// fmt package Scanf function
	fmt.Println("What's your age?")
	var age string
	_, err = fmt.Scanf("%s", &age)
	if err != nil {
		return
	}
	fmt.Println("Hello " + name + " " + name2 + " " + age + " years old")
}
