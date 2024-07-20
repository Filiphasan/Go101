package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Type Cast
	var a, b = 17, 45
	var div = a / b
	fmt.Println("Type Cast int: ", div)
	var divFloat = float64(a) / float64(b)
	fmt.Println("Type Cast float: ", divFloat)

	// string conversion
	var str = "10"
	var num, _ = strconv.Atoi(str)
	fmt.Println("Type Cast string: ", num)

	// string conversion
	var str2 = "10.5"
	var num2, _ = strconv.ParseFloat(str2, 64)
	fmt.Println("Type Cast string: ", num2)

	// string conversion
	var str3 = "true"
	var num3, _ = strconv.ParseBool(str3)
	fmt.Println("Type Cast string: ", num3)

	// restring conversion
	var num4 = 10
	var str4 = strconv.Itoa(num4)
	fmt.Println("Type Cast string: ", str4)

	// restring conversion
	var num5 = 10.5
	var str5 = strconv.FormatFloat(num5, 'f', 2, 64)
	fmt.Println("Type Cast string: ", str5)

	// restring conversion
	var num6 = true
	var str6 = strconv.FormatBool(num6)
	fmt.Println("Type Cast string: ", str6)
}
