package main

import (
	"fmt"
	"net/http"
)

func welcomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Store!"))
}

func main() {
	http.HandleFunc("/welcome", welcomePage)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
