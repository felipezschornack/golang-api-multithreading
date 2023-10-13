package util

import (
	"encoding/json"
	"log"
)

func PrintDataAsJson(apiName string, data any) {
	formattedData, err := json.MarshalIndent(data, "", "   ")
	checkError(err)
	log.Printf("\nData obtained using %s\n%s\n", apiName, string(formattedData))
}

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
