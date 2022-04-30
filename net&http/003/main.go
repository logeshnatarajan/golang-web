package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	switch req.URL.Path {
	case "/dog":
		fmt.Fprintln(w, "dog dog doggyyyyy")
	case "/cat":
		fmt.Fprintln(w, "cat cat catty")

	}

}

func main() {
	var m hotdog
	http.ListenAndServe(":8080", m)
}
