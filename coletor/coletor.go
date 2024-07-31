package coletor

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"time"
)

type Coleta struct {
	Titulo    string `json:"titulo"`
	Descricao string `json:"desc"`
}

// Função para fazer scrapping no DuckDuckGo
// Recebe uma string para ser pesquisada
// Devolve um objeto Coleta com os campo descrição e contexto.
func ColetarDados(descricao string) ([]Coleta, error) {
	dadosColetados := make([]Coleta, 0)
	fmt.Println("Pesquisa:", descricao)

	sliceTitulos := make([]string, 0)
	sliceDesc := make([]string, 0)

	// Instância do coletor de dados
	coletor := colly.NewCollector()

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
	for i := 0; i < 3; i++ { // Tenta até 3 vezes
		err := coletor.Visit(searchURL)
		if err == nil {
			break
		}
		fmt.Printf("Retentativa %d/3\n", i+1)
		time.Sleep(1 * time.Second)
	}

	coletor.Wait()

	for indice, valor := range sliceTitulos {
		dadosColetados = append(dadosColetados, Coleta{valor, sliceDesc[indice]})
	}
	return dadosColetados, nil
}
