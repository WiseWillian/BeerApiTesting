package beers_models

// Autor: Rafael Willian
//
// Descrição: Uma biblioteca que possui o modelo de descrição do endpoint /beers em estruturas de dados,
// assim como a função que caracteriza a criação de objetos a partir destes modelos

import (
	"encoding/json"
)

//Guarda a categoria do estilo
type Category struct {
	Id int `json:"id"`
	Name string `json:"name"`
	CreateDate string `json:"createDate"`
}

//Guarda detalhes sobre a disponibilidade da bebida
type Available struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

//Guarda a informação sobre o estilo da bebida
type Style struct {
	Id int `json:"id"`
	CategoryId int `json:"categoryId"`
	StyleCategory Category `json:"category"`
	Name string `json:"name"`
	ShortName string `json:"shortName"`
	Description string `json:"description"`
	IbuMin string `json:"ibuMin"`
	IbuMax string `json:"ibuMax"`
	AbvMin string `json:"abvMin"`
	AbvMax string `json:"abvMax"`
	SrmMin string `json:"srmMin"`
	SrmMax string `json:"srmMax"`
	OgMin string `json:"ogMin"`
	FgMin string `json:"fgMin"`
	FgMax string `json:"fgMax"`
	CreateDate string `json:"createDate"`
	UpdateDate string `json:"updateDate"`
}

//Guarda a informação da bebida individualmente
type Beer struct {
	Id string `json:"id"`
	Name string `json:"name"`
	NameDisplay string `json:"nameDisplay"`
	Description string `json:"description"`
	Abv string `json:"abv"`
	Ibu string `json:"ibu"`
	AvailableId int `json:"availableId"`
	StyleId int `json:"styleId"`
	IsOrganic string `json:"isOrganic"`
	Status string `json:"status"`
	StatusDisplay string `json:"statusDisplay"`
	CreateDate string `json:"createDate"`
	UpdateDate string `json:"updateDate"`
	BeerAvailable Available `json:"available"`
	BeerStyle Style `json:"style"`
}

//Guarda toda a informação da requisição
type BeerApiResponse struct {
	CurrentPage int `json:"currentPage"`
	NumberOfPages int `json:"numberOfPages"`
	TotalResults int `json:"totalResults"`
	Data []Beer `json:"data"`
}

//Função que transforma o array json em dados estruturados
func GetAllBeers(body []byte) (*BeerApiResponse, error){
	var beers = new (BeerApiResponse)
	err := json.Unmarshal(body, &beers)

	return beers, err
}