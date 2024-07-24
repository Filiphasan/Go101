package main

import "fmt"

type Kilometer float64

func (km Kilometer) toFeet() float64 {
	return float64(km) * 3280.84
}

func main() {
	fmt.Println("Enter the distance in kilometers")
	var distance Kilometer
	_, err := fmt.Scan(&distance)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(distance.toFeet())
}
