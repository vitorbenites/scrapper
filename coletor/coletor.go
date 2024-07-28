package coletor

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

// Formato da resposta
type Coleta struct {
	Desc     string `json:"desc"`
	Contexto string `json:"contexto"`
}

// Função para fazer scrapping no DuckDuckGo
// Recebe uma string para ser pesquisada
// Devolve um objeto Coleta com os campo descrição e contexto.
func ColetarDados(descricao string) Coleta {
	var dadosColetados Coleta
	dadosColetados.Desc = descricao
	fmt.Println("Pesquisa:", descricao)

	sliceTitulos := make([]string, 10)
	sliceDesc := make([]string, 10)

	// Instância do coletor de dados
	coletor := colly.NewCollector(
		colly.AllowedDomains("duckduckgo.com", "html.duckduckgo.com"),
	)

	// Callback para quando um elemento com a tag <a> for encontrado
	coletor.OnHTML("a.result__a", func(e *colly.HTMLElement) {
		sliceTitulos = append(sliceTitulos, e.Text)
	})

	// Callback para quando um elemento com a classe result__snippet for encontrado
	coletor.OnHTML("a.result__snippet", func(e *colly.HTMLElement) {
		sliceDesc = append(sliceDesc, e.Text)
	})

	// Callback para quando a solicitação terminar
	coletor.OnScraped(func(r *colly.Response) {
		fmt.Println("Raspagem concluída.")
	})

	// Realiza a busca
	searchQuery := descricao
	searchURL := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", searchQuery)
	err := coletor.Visit(searchURL)
	if err != nil {
		log.Fatal(err)
	}

	for indice, valor := range sliceTitulos {
		dadosColetados.Contexto += "Título: " + valor + "\n"
		dadosColetados.Contexto += "Descrição: " + sliceDesc[indice] + "\n"
	}
	return dadosColetados
}
