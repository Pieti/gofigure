package main

import (
	"github.com/pieti/gofigure/credit_card_validator/luhn"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type CreditCard struct {
	Number string `json:"number"`
	Name   string `json:"name"`
	Expiry string `json:"expiry"`
	Valid  bool   `json:"valid"`
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /{$}", index)
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	cardNumber := r.FormValue("card_number")
	valid := luhn.Validate(cardNumber)

	card := CreditCard{
		Number: cardNumber,
		Name:   r.FormValue("card_holder"),
		Expiry: r.FormValue("expiry"),
		Valid:  valid,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", card)
}
