package main

import "fmt"

func main() {
	// int format
	var a = 10
	fmt.Printf("This is a int: %d\n", a)

	// string format
	var b = "hello"
	fmt.Printf("This is a string: %s\n", b)

	// float format
	var c = 3.14
	fmt.Printf("This is a float: %f\n", c)

	// float format with precision
	var d = 3.1415926
	fmt.Printf("This is a float with precision: %.2f\n", d)

	// string format with width
	var e = "hello"
	fmt.Printf("This is a string with width: %s\n", e)

	// string format with width and precision
	var f = "hello"
	fmt.Printf("This is a string with width and precision: %5.2s\n", f)

	// string format with Sprintf function
	var firstName = "Hasan"
	var lastName = "Erdal"
	var fullName = fmt.Sprintf("%s %s", firstName, lastName)
	fmt.Println(fullName)
}
