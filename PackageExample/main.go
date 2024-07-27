package main

import (
	"fmt"
	"github.com/sethvargo/go-password/password"
)

func main() {
	res, err := password.Generate(16, 4, 4, false, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
