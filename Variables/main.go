package main

import (
	"fmt"
)

func main() {
	// string type for text
	var message string = "deneme"
	fmt.Println(message)

	// int type for whole numbers
	var firstNumber int = 1234
	fmt.Println(firstNumber)

	// float type for decimal numbers
	var decimalNumber float32 = 3.14
	fmt.Println(decimalNumber)

	// bool type for true or false
	var isTrue bool = true
	fmt.Println(isTrue)

	// rune type for unicode characters, char in csharp or java ...
	var letter rune = 'a'
	fmt.Println(letter)

	// byte type for ascii characters
	var letter2 byte = 12
	fmt.Println(letter2)

	// multiple variables
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// short variable declaration
	d, e, f := 4, 5, 6
	fmt.Println(d, e, f)

	// multiple variable declaration
	var g, h, i = 7, 8.23, "Hasan"
	fmt.Println(g, h, i)

	// declare and set variable
	var j int
	j = 9
	fmt.Println(j)

	// declare, print after set variable
	var k = 10
	fmt.Println(k)
	k = 33
	fmt.Println(k)

	// constant variable
	const pi = 3.14
	fmt.Println(pi)
}
