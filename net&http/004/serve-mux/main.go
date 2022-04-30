package main

import (
	"fmt"
	"net/http"
)

type hotdog int
type hotcat int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "dog dog ddoggyyy")

}
func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "cat cat cattyyy")

}

func main() {
	var m hotdog
	var g hotcat
	mux := http.NewServeMux()
	mux.Handle("/dog", m)
	mux.Handle("/cat", g)

	http.ListenAndServe(":8080", mux)

}
