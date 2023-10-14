package viacep

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// For more information, please see https://viacep.com.br
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func GetZipCode(zipCode string, channel chan<- ViaCEP) {
	url := "https://viacep.com.br/ws/" + zipCode + "/json/"
	log.Printf("Searching zip code [%s] using ViaCEP [%s]", zipCode, url)

	resp, err := http.Get(url)
	checkError("HTTP request error [%s]: %v\n", url, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkError("Error reading ResponseBody [%s]: %v\n", url, err)

	var data ViaCEP
	err = json.Unmarshal(body, &data)
	checkError("Error converting Json to object [%s]: %v\n", url, err)

	channel <- data
}

func checkError(message, url string, err error) {
	if err != nil {
		log.Panicf(message, url, err)
	}
}
