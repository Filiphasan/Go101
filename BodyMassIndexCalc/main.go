package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	weight := 70
	height := 1.70
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your weight")
	scanner.Scan()
	weight, _ = strconv.Atoi(scanner.Text())
	fmt.Println("Enter your height")
	scanner.Scan()
	height, _ = strconv.ParseFloat(scanner.Text(), 64)
	fmt.Println("Your BMI is:", CalculateBmi(weight, height))

}

func CalculateBmi(weight int, height float64) float64 {
	return float64(weight) / (height * height)
}
