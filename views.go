package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to the Users API")
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Users")
	w.Header().Set("Content-Type", "application/json")
	rows := sqlQuery("select id, name, email from users")
	defer rows.Close()
	var (
		id    int
		name  string
		email string
		user  User
		users []User
	)
	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		checkError(err)
		user = User{UserId: id, Name: name, Email: email}
		users = append(users, user)
	}
	if len(users) == 0 {
		json.NewEncoder(w).Encode("no users present currently")
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create User")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("No Data Passed")
		return
	}
	// whatif: only {} is provided
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	if user.isEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}
	err := user.addToUsers()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode("User Added")
	}
}
