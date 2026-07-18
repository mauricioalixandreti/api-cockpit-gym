package main

import (
	"API-COCKPIT-GYM/repository"
	"log"
	"net/http"
)

func main() {
	// 1. Inicializa o Firebase que está na pasta Repository
	repository.InitFirebase()

	// 2. Configura o Go para servir os arquivos estáticos da sua pasta View
	fs := http.FileServer(http.Dir("./View/EndPoints"))
	http.Handle("/", fs)

	log.Println("Servidor web rodando em http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
