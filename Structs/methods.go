package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	notes     []int
	isPass    bool
}

func (s Student) getFullName() string {
	return s.firstName + " " + s.lastName
}

func (s Student) getAverage() float64 {
	sum := 0
	for _, note := range s.notes {
		sum += note
	}
	return float64(sum) / float64(len(s.notes))
}

func main() {

	var student = Student{
		firstName: "Hasan",
		lastName:  "Erdal",
		notes:     []int{35, 80},
		isPass:    true,
	}
	fmt.Println(student.getFullName())
	fmt.Println(student.getAverage())
}
