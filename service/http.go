package service

import (
	"fmt"
	"net/http"
)

// Run function
func Run() error {
	fmt.Println("Running on 8080")
	http.HandleFunc("/deploy/", deploy)
	http.HandleFunc("/", home)
	return http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func deploy(w http.ResponseWriter, r *http.Request) {
	start := len("/deploy/")
	fmt.Fprintf(w, Deploy(r.URL.Path[start:]))
}
