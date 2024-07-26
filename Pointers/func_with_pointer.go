package main

import "fmt"

type User struct {
	firstName string
	lastName  string
}

func main() {

	number := 10
	increment(&number)
	fmt.Println(number)

	user := User{"Hasan", "Erdal"}
	pUser := &user
	pUser.updateFirstName("Mustafa")
	fmt.Println(user.firstName)

}

func increment(number *int) {
	*number += 1
}

func (u *User) updateFirstName(firstName string) {
	u.firstName = firstName
}
