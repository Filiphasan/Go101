package main

import "fmt"

func main() {

	var vize int
	var final int
	var examScore int
	fmt.Println("What's your vize score?")
	_, err := fmt.Scan(&vize)
	if err != nil {
		fmt.Println("This is not a number", err.Error())
		return
	}
	fmt.Println("What's your final score?")
	_, err = fmt.Scan(&final)
	if err != nil {
		fmt.Println("This is not a number", err.Error())
		return
	}
	examScore = calcExamScore(vize, final)
	fmt.Println("Your exam score is:", examScore)
}

func calcExamScore(vize int, final int) int {
	vizeScore := vize * 4 / 10
	finalScore := final * 6 / 10
	return vizeScore + finalScore
}
