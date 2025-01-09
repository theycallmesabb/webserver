package main

import (
	"fmt"
	"log"
	"net/http"
)

// Form handler for POST request processing
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post Request Successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s \n", name)
	fmt.Fprintf(w, "Address = %s \n", address)
}

// Hello handler for custom route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func main() {
	// Serve static files from the "Static" directory (correct path)
	fileServer := http.FileServer(http.Dir("./Static"))
	http.Handle("/", http.StripPrefix("/", fileServer)) // Serve static files from Static/

	// Form submission handler
	http.HandleFunc("/form", formHandler)
	// Custom hello route
	http.HandleFunc("/hello", helloHandler)

	// Start the server
	fmt.Println("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
