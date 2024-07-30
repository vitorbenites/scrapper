package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/vitorbenites/scrapper/gerenciador"
)

func main() {
	runtime.GOMAXPROCS(4)
	// Função Handle
	http.HandleFunc("/", gerenciador.GerenciarRequisicao)
	// Inicialização do servidor
	fmt.Println("Servidor escutando na porta 5000")
	http.ListenAndServe(":5000", nil)
}
