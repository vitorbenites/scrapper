package gerenciador

import (
	"encoding/json"
	"github.com/vitorbenites/scrapper/coletor"
	"io"
	"net/http"
)

// Formato da requisição
type Requisicao struct {
	Desc string `json:"desc"`
}

// Função ReceberRequisicao para receber e responder requisições
func GerenciarRequisicao(writer http.ResponseWriter, reqRecebida *http.Request) {
	// Verificar se o método HTTP é POST
	if reqRecebida.Method != http.MethodPost {
		http.Error(writer, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Ler o corpoReq da requisição
	corpoReq, err := io.ReadAll(reqRecebida.Body)
	if err != nil {
		http.Error(writer, "Erro ao ler o corpo da requisição.", http.StatusBadRequest)
	}

	// Decodificar o json da requisição
	var requisicao Requisicao
	err = json.Unmarshal(corpoReq, &requisicao)
	if err != nil {
		http.Error(writer, "Erro ao decodificar JSON.", http.StatusBadRequest)
		return
	}

	// Processamento dos dados da requisição
	dadosColetados := coletor.ColetarDados(requisicao.Desc)

	// Configuração do cabeçalho da resposta para JSON
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	// Codificar a resposta em JSON
	respostaJSON, err := json.Marshal(dadosColetados)
	if err != nil {
		http.Error(writer, "Erro ao codificar JSON", http.StatusInternalServerError)
		return
	}

	// Envio da resposta
	writer.Write(respostaJSON)
}
