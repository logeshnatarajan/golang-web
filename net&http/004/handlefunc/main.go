package main

import (
	"fmt"
	"net/http"
)

func m(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "dog dog ddoggyyy")

}
func g(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "cat cat cattyyy")

}

func main() {

	http.HandleFunc("/dog/", m)
	http.HandleFunc("/cat", g)

	http.ListenAndServe(":8080", nil)

}
