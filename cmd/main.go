package main

import (
	"log"
	"os"
	"time"

	"github.com/felipezschornack/multithreading/api/apicep"
	"github.com/felipezschornack/multithreading/api/viacep"
	"github.com/felipezschornack/multithreading/util"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Please, enter a zip code using one of the following formats: 00000-000 or 00000000")
	} else {
		zipCode, err := util.FormatZipCode(os.Args[1])
		checkError(err)
		getZipCode(zipCode)
	}
}

func getZipCode(zipCode string) {
	channelApiCEP := startApiCEP(zipCode)
	channelViaCEP := startViaCEP(zipCode)

	select {
	case result := <-channelApiCEP:
		util.PrintDataAsJson("ApiCEP", result)
	case result := <-channelViaCEP:
		util.PrintDataAsJson("ViaCEP", result)
	case <-time.After(time.Second):
		log.Println("Timeout")
	}
}

func startApiCEP(zipCode string) chan apicep.ApiCEP {
	channelApiCEP := make(chan apicep.ApiCEP)
	go apicep.GetZipCode(zipCode, channelApiCEP)
	return channelApiCEP
}

func startViaCEP(zipCode string) chan viacep.ViaCEP {
	channelViaCEP := make(chan viacep.ViaCEP)
	go viacep.GetZipCode(zipCode, channelViaCEP)
	return channelViaCEP
}

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
