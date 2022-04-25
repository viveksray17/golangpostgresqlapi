package main

import (
	"fmt"
)

type User struct {
	UserId int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func (u User) addToUsers() error {
	err := sqlExec(fmt.Sprintf("insert into users(name, email) values('%v', '%v');", u.Name, u.Email))
	return err
}
func (u User) dropFromUsers() {
	sqlExec(fmt.Sprintf("delete from users where email = '%v'", u.Email))
}

func (u User) isEmpty() bool {
	return u.Name == "" && u.Email == ""
}
