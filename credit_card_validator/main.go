package main

import (
	"github.com/pieti/gofigure/credit_card_validator/luhn"
	"html/template"
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
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
		return
	}

	if r.Method == http.MethodPost {
		cardNumber := r.FormValue("card_number")
		valid := luhn.Validate(cardNumber)

		card := CreditCard{
			Number: cardNumber,
			Name:   r.FormValue("card_holder"),
			Expiry: r.FormValue("expiry"),
			Valid:  valid,
		}
		tpl.ExecuteTemplate(w, "index.gohtml", card)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
