//
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler) // register handler to web page
	http.ListenAndServe(":8080", nil)
}
