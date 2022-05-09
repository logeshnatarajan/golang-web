package main

import (
	"fmt"
	"net/http"
)

type hotdog int
// type hotdog is attached to this function
func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "any code u want in this func")

}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)

}
