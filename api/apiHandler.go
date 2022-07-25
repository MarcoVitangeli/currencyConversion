package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type CurrencyApiHandler struct {
	From string
	To   string
}

func (apiHandler *CurrencyApiHandler) buildUrl() string {
	return fmt.Sprintf("https://currency-exchange.p.rapidapi.com/exchange?to=%s&from=%s&q=1.0", apiHandler.To, apiHandler.From)
}

func (apiHandler *CurrencyApiHandler) GetRawData() float32 {
	url := apiHandler.buildUrl()
	req, err := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("X-RapidAPI-Key", "a8fe49be17msh04f20561e45423bp191780jsn2c62f34aff4d")
	req.Header.Add("X-RapidAPI-Host", "currency-exchange.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Panic(err)
	}

	result, err := strconv.ParseFloat(string(body), 32)

	if err != nil {
		log.Panic(err)
	}
	return float32(result)
}
