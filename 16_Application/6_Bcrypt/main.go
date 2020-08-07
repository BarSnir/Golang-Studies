package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	p := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	login := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(login))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Login Successfully")
	}
}
