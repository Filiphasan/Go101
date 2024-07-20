package main

import "fmt"

func main() {
	// Arithmetic operators
	var a = 10
	var b = 20
	var addition = a + b
	var minus = b - a
	var multiply = a * b
	var divide = b / a
	var mod = b % a
	fmt.Println("Arithmetic operators", a, b, addition, minus, multiply, divide, mod)

	// Assignment operators
	var x = 11
	var y = 23
	var z = 5
	x += z
	fmt.Println("Assignment operators", x, y)
	y -= z
	fmt.Println("Assignment operators", x, y)
	y /= z
	fmt.Println("Assignment operators", x, y)
	x *= z
	fmt.Println("Assignment operators", x, y)
	x %= z
	fmt.Println("Assignment operators", x, y)

	// Relational operators
	var equal = a == b
	var notEqual = a != b
	var greater = a > b
	var less = a < b
	var greaterOrEqual = a >= b
	var lessOrEqual = a <= b
	fmt.Println("Relational operators", equal, notEqual, greater, less, greaterOrEqual, lessOrEqual)

	// Logical operators
	var and = a > 10 && b > 10
	var or = a > 10 || b > 10
	var not = !(a > 10)
	fmt.Println("Logical operators", and, or, not)
}
