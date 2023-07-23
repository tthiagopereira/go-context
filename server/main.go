package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3005", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Iniciou minha request")
	defer log.Println("Finalizado a request")

	select {
	case <-time.After(time.Second * 5):
		fmt.Fprintln(w, "Requisicao processada com sucesso")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Requicao processada com sucesso"))
	case <-ctx.Done():
		log.Println("Request cancelada")
		http.Error(w, "Request cancelada", http.StatusRequestTimeout)
	}
}
