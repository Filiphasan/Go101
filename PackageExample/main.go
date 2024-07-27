package main

import (
	"Go101/PackageExample/customPackage"
	"fmt"
	"github.com/sethvargo/go-password/password"
)

func main() {
	res, err := password.Generate(16, 6, 4, false, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	customPackage.HelloCustomPackage()
}
