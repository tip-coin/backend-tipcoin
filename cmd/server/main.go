package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hi, i'm alive tip coin bruh!")
	})

	fmt.Println("server listening on: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
