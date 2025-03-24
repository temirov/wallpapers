package main

import (
	"fmt"
	"net/http"
)

func main() {
	staticFileServer := http.FileServer(http.Dir("."))
	http.Handle("/", staticFileServer)

	fmt.Println("Server is running on http://localhost:8080")
	errorValue := http.ListenAndServe(":8080", nil)
	if errorValue != nil {
		panic(errorValue)
	}
}
