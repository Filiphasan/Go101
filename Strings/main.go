package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "HASAN ERDAL"
	name2 := "hasan erdal"
	fmt.Println(strings.ToLower(name))  // lower
	fmt.Println(strings.ToUpper(name2)) // upper

	fmt.Println(strings.Index(name, "ERDAL"))                  // index
	fmt.Println(strings.Count(name, "A"))                      // count
	fmt.Println(strings.Contains(name, "ERDAL"))               // true or false
	fmt.Println(strings.Compare(strings.ToLower(name), name2)) // 0 for equal, 1 for bigger, -1 for smaller
	fmt.Println(strings.Split(name, " "))                      // split
	fmt.Println(strings.Replace(name, "ERDAL", "Erdal", 1))    // replace with count
	fmt.Println(strings.ReplaceAll(name, "ERDAL", "Erdal"))    // replace all
}
