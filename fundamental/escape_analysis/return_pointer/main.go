package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func newUser(id int, name string) *User {
	u := User{ID: id, Name: name} // u створюється локально
	return &u                     // повертаємо вказівник на локальну змінну u
}

func main() {
	userPtr := newUser(1, "Vladyslav")
	fmt.Println("User ID:", userPtr.ID)
}

// команда компіляції:
// go build -gcflags="-m" ./fundamental/escape_analysis/return_pointer/main.go
