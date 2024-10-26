package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Post req successfull\n")

	// Parse the form
	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Email: %s\n", email)

	// DO YOUR LOGIC WITH THE FORM VALUES
}

func main() {
	// telling golang that you have to check out the static folder
	// automatically looks into the index html file
	fileServer := http.FileServer(http.Dir("./static")) // Here "." is important before static path else it will try to check on the root path

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is listening on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
