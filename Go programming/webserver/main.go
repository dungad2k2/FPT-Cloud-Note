package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Noooooob!!!. Request Path: %s", r.URL.Path)
}
type User struct{
	Name string `json:"username"`
	Email string `json:"email"`
	Age int `json:"age"`
	IsAdmin bool
}
func UserHandler(w http.ResponseWriter, r *http.Request){
	user := User{
		Name:    "Gopher",
		Email:   "gopher@google.com",
		Age:     10,
		IsAdmin: true,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func Add(x int, y int) int {
	return x + y
}
func main(){
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api/user", UserHandler)
	fmt.Println("Starting server on port 8080....")
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}