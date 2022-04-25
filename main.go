package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	UserId int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func sqlExec(query string) error {
	psqlConnection := fmt.Sprintf("host=localhost port=5432 user=vivek password=%v dbname=vivek sslmode=disable", os.Getenv("PG_PASS"))
	db, _ := sql.Open("postgres", psqlConnection)
	defer db.Close()
	_, err := db.Exec(query)
	return err
}

func sqlQuery(query string) *sql.Rows {
	psqlConnection := fmt.Sprintf("host=localhost port=5432 user=vivek password=%v dbname=vivek sslmode=disable", os.Getenv("PG_PASS"))
	db, err := sql.Open("postgres", psqlConnection)
	defer db.Close()
	checkError(err)
	rows, err := db.Query(query)
	checkError(err)
	return rows
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

func main() {
	fmt.Println("Users API")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/users", getAllUsers).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":4000", r))
}

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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
