package main

import (
	"fmt"
	"io"
	"net/http"
)

type Teacher struct {
	Id               string `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	RegistrationDate string `json:"registration_date"`
}

type Student struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Group struct {
	Id string `json:"id"`
}

type Subject struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Teacher Teacher `json:"teacher"`
}

func main() {
	http.HandleFunc("/", getRoot)     // Website
	http.HandleFunc("/api", getHello) //

	err := http.ListenAndServe(":8080", nil)
	panic(err.Error())
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
