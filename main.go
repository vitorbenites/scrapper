package main

import (
	"fmt"
	"github.com/vitorbenites/scrapper/gerenciador"
	"net/http"
)

func main() {
	// Função Handle
	http.HandleFunc("/", gerenciador.GerenciarRequisicao)
	// Inicialização do servidor
	fmt.Println("Servidor escutando na porta 5000")
	http.ListenAndServe(":5000", nil)
}
