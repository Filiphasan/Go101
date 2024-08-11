package main

import (
	"fmt"
	"os"
)

func main() {

	// Create a new file
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
	fmt.Println("File created successfully!")

	// Open
	fileOpened, err := os.OpenFile("file.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write
	_, err = fileOpened.WriteString("Hello, World!")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File written successfully!")
}
