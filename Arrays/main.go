package main

import "fmt"

func main() {
	// Array integer
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Basic Array:", numbers)

	// Array string
	var names = [4]string{"John", "Jane", "Joe", "?"}
	fmt.Println("String Array:", names)
	names[3] = "Hasan"
	fmt.Println("String Array:", names)

	// Array dynamic
	var numbers2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array dynamic:", numbers2)

	// multi dimensional array
	var numbers6 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Multi dimensional array", numbers6)

	// Array len function
	var numbers3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array length:", len(numbers3))

	// Array cap function
	var numbers4 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array capacity:", cap(numbers4))

	// slice of array
	var numbers5 = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice of array", numbers5[2:4])
	var slice = numbers5[2:4]
	fmt.Println("Slice of array length", len(slice))
	fmt.Println("Slice of array capacity", cap(slice))

	// slice append
	slice = append(slice, 6)
	slice = append(slice, 7)
	fmt.Println("Slice of array", slice)
	fmt.Println("Slice of array length", len(slice))
	fmt.Println("Slice of array capacity", cap(slice))

	// slice make
	slice2 := make([]int, 5, 10)
	fmt.Println("Slice of array", slice2)
	fmt.Println("Slice of array length", len(slice2))
	fmt.Println("Slice of array capacity", cap(slice2))

	// array with range
	var numbers7 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array with range")
	for i, value := range numbers7 {
		fmt.Println(i, value)
	}

	// maps - standart
	var students = map[int]string{1: "John", 2: "Jane", 3: "Joe"}
	fmt.Println("Maps:", students)
	fmt.Println("Maps with range")
	for key, value := range students {
		fmt.Println(key, value)
	}

	// maps - dynamic
	var titles = make(map[string]string)
	titles["Dr"] = "Doctor"
	titles["Mr"] = "Master"
	fmt.Println("Maps:", titles)
	fmt.Println("Maps with range")
	for key, value := range titles {
		fmt.Println(key, value)
	}
}
