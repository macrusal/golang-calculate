package app

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bem vindo a API de calculos"))
}

func StartApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Iniciando o servidor na porta: 4010")
	err := http.ListenAndServe(":4010", mux)
	log.Fatal(err)
}
