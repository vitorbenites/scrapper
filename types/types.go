package types

type Coleta struct {
	Titulo    string `json:"titulo"`
	Descricao string `json:"desc"`
}

// Formato da requisição
type Requisicao struct {
	Desc string `json:"desc"`
}
