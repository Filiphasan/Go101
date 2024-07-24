package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	Email     string
	birthDate time.Time
	account   UserAccount
}

type UserAccount struct {
	userName  string
	password  string
	loginDate time.Time
}

func main() {

	var user = User{
		firstName: "Hasan",
		lastName:  "Erdal",
		Email:     "h@h.com",
		birthDate: time.Date(1998, 6, 17, 0, 0, 0, 0, time.UTC),
		account: UserAccount{
			userName:  "hasan.erdal",
			password:  "12345",
			loginDate: time.Now(),
		},
	}
	fmt.Println("User:", user)

	// anonymous struct
	admin := struct {
		user    User
		isAdmin bool
	}{
		user:    user,
		isAdmin: true,
	}
	fmt.Println("Admin:", admin)

}
