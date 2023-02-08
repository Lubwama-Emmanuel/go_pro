package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, World!")
		if err != nil {
			fmt.Println("An error occured", err)
		}
		fmt.Println("Here are the bytes:", n)
	})

	_ = http.ListenAndServe(":8080", nil)
}
