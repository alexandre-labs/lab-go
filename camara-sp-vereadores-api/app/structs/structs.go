package structs

type Vereadores struct {
	Id  string
	Url []string
}

type Vereador struct {
	Id      string
	Url     []string
	RawData *RawData
}

type RawData struct {
	Nome     string
	Contato  *Contato
	Projetos []Projeto
	Partido  *Partido
}

type Contato struct {
	Fax             string
	Internet        *Internet
	Telefone        string
	Correspondencia *Correspondencia
}

type Internet struct {
	Email string
}

type Correspondencia struct {
	Sala  string
	Andar string
}

type Projeto struct {
	Url    string
	Nome   string
	Status string
}

type ProjetosQtde struct {
	aprovadorQtde  int
	vetadoQtde     int
	tramitacaoQtde int
}

type Partido struct {
	Image string
	Sigla string
	Nome  string
}
