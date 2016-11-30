# camara-sp-vereadores-api

API simples reproduzindo a API que oferece dados sobre vereadores da Câmara Municipal de São Paulo

## Diretório data

Nesse diretório estão os dados recuperados através do endpoint: "http://www1.camara.sp.gov.br/vereador_json.asp"

O código responsavel por obter os IDs dos vereadores e baixar o conteúdo do endpoint citado está dísponivel [nesse repositório](https://github.com/cdmb/camara-sp-vereadores/tree/development).

## Banco de dados

Os dados disponiveis no diretório `data` são persistidos em uma base de dados `MongoDB` utilizando o script `store-data.go`:

```bash
go run store-data.go
```

## Endpoints

Nesse experimento há apenas dois endpoints:

### /vereadores

Retorna dados simples de todos os vereadores disponiveis no diretório `data`.

#### Response

```json
{
    "Id": "string",
    "Url": "[]string"
}
```

### /vereador/:id

Retorna detalhes de um vereador específico. O parametro id é o `Id` recuperado pelo endpoint
`/vereadores`.

#### Response

Retorna a estrutura [Vereador](https://github.com/alexandre/lab-go/blob/master/camara-sp-vereadores-api/app/structs/structs.go#L8).

## Instalação

* Clonar o repositório:

```bash
git clone https://github.com/alexandre/lab-go ${GOPATH}/src/github.com/alexandre/lab-go
```

* Instalar dependências:

```
go get github.com/julienschmidt/httprouter

go get go get gopkg.in/mgo.v2
```

* Executar app:

```
cd ${GOPATH}/src/github.com/alexandre/lab-go/camara-sp-vereadores-api

go run app/app.go
```

## TODO

- [ ] Escrever testes
- [ ] Filtrar vereadores utillizando querystrings (e.g.: `?partido=PCdoB`)
