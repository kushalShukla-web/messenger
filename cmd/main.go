package main

import (
	"fmt"
	"net/http"
)

func kushal(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello from Kushal side!")
}
func main() {

	http.HandleFunc("/handle", kushal)
	http.ListenAndServe(":8080", nil)
}
