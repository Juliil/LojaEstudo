package main

import (
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Preta", Preco: 39, Quantidade: 1},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 189, 1},
		{"Filhote de Miu", "Gatão", 19999, 1},
	}

	temp.ExecuteTemplate(w, "Index", produtos)

}
