package main

import (
	"html/template"
	"log"
	"os"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("abc.get"))

}
func main() {
	sages := map[string]string{
		"india ": "gandhi",
		"africa": "mandala",
	}

	err := t.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatal(err)
	}

}
