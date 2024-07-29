package coletor

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
)

// Função para fazer scrapping no DuckDuckGo
// Recebe uma string para ser pesquisada
// Devolve um objeto Coleta com os campo descrição e contexto.
func ColetarDados(descricao string) (map[string]string, error) {
	dadosColetados := make(map[string]string)
	fmt.Println("Pesquisa:", descricao)

	sliceTitulos := make([]string, 0)
	sliceDesc := make([]string, 0)

	// Instância do coletor de dados
	coletor := colly.NewCollector(
	// colly.AllowedDomains("html.duckduckgo.com"),
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

	coletor.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Algo deu errado:", err)
	})

	// Realiza a busca
	searchQuery := descricao
	searchURL := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s&kl=br-pt", url.QueryEscape(searchQuery))
	err := coletor.Visit(searchURL)
	if err != nil {
		return nil, fmt.Errorf("Erro ao visitar URL: %v", err)
	}

	coletor.Wait()

	for indice, valor := range sliceDesc {
		dadosColetados[sliceTitulos[indice]] = valor
	}
	return dadosColetados, nil
}
