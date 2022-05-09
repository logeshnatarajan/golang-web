package main

import (
	"fmt"
	"net/http"
)

type hotdog int
// type hotdog is a method of this function
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("logesh", "podu podu varuthu")
	w.Header().Set("logesh", "enna da varuthuuuu")
	fmt.Fprint(w, "money money moneyyyy")
}

func main() {
	var typehotdog hotdog
	http.ListenAndServe(":8080", typehotdog)

}
