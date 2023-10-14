package apicep

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// For more information, please see https://apicep.com/api-de-consulta
type ApiCEP struct {
	Status   int    `json:"status"`
	Code     string `json:"code"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}

func GetZipCode(zipCode string, channel chan<- ApiCEP) {
	url := "https://cdn.apicep.com/file/apicep/" + zipCode + ".json"
	log.Printf("Searching zip code [%s] using ApiCEP [%s]", zipCode, url)

	resp, err := http.Get(url)
	checkError("HTTP request error [%s]: %v\n", url, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkError("Error reading ResponseBody [%s]: %v\n", url, err)

	var data ApiCEP
	err = json.Unmarshal(body, &data)
	checkError("Error converting Json to object [%s]: %v\n", url, err)

	channel <- data
}

func checkError(message, url string, err error) {
	if err != nil {
		log.Panicf(message, url, err)
	}
}
