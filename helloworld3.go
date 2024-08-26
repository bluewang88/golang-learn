package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World from w.Write"))
	fmt.Fprintf(w, "\nHello, World from fmt.Fprintf! URL : %s", r.URL.Path)
}

// Run the server
// go run helloworld3.go
// Open the browser and go to http://localhost:8080
// You should see "Hello, World!" in the browser.
