package main

import (
	"github.com/currency/api"
	"github.com/currency/email"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	apiHandler := api.CurrencyApiHandler{
		To:   "ARS",
		From: "USD",
	}

	currData := apiHandler.GetRawData()
	godotenv.Load(".env")

	emailHand := email.EmailHandler{
		From: os.Getenv("ADDR"),
		// TODO: send to more contacts
		To:       []string{os.Getenv("TO")},
		Password: os.Getenv("PASS"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
	}

	emailHand.SendCurrInfo(currData)
}
