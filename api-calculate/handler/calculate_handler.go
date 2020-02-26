package handler

import (
	"fmt"
	"net/http"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chegou no handler...")

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("O que tem no r..: ", r.URL)
	fmt.Println("O que tem no r..: ", r.URL.Path)
	fmt.Println("O que tem no r..: ", r.URL.Path[len("/calculate/"):])
	fmt.Println("O que tem no r..: ", r.URL.Path[len("/calculate/"):])
}
