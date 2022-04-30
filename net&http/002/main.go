package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("logesh", "podu podu varuthu")
	w.Header().Set("logesh", "enna da varuthuuuu")
	fmt.Fprint(w, "money money moneyyyy")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)

}
