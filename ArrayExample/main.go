package main

import "fmt"

func main() {
	var firstArrLen = 10
	var secondArrLen = 10
	fmt.Println("Enter the length of first array")
	_, err := fmt.Scan(&firstArrLen)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Enter the length of second array")
	_, err2 := fmt.Scan(&secondArrLen)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println("Enter the elements of first array")
	var firstArr = make([]int, firstArrLen)
	for i := 0; i < firstArrLen; i++ {
		_, err3 := fmt.Scan(&firstArr[i])
		if err3 != nil {
			fmt.Println(err3)
			return
		}
	}

	fmt.Println("Enter the elements of second array")
	var secondArr = make([]int, secondArrLen)
	for i := 0; i < secondArrLen; i++ {
		_, err4 := fmt.Scan(&secondArr[i])
		if err4 != nil {
			fmt.Println(err4)
			return
		}
	}

	fmt.Println("Matrix of arrays")

	for i := 0; i < firstArrLen; i++ {
		for j := 0; j < secondArrLen; j++ {
			fmt.Print(firstArr[i]*secondArr[j], " ")
		}
		fmt.Println()
	}
}
