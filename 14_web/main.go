package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("SERVER STARTED AT - PORT 3000")
	http.ListenAndServe(":3000", nil)
}
