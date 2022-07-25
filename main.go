package main

import (
	"github.com/currency/api"
	"github.com/currency/contact"
	"github.com/currency/email"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	apiHandler := api.CurrencyApiHandler{
		To:   "ARS",
		From: "USD",
	}

	currData := apiHandler.GetRawData()
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic(err)
	}

	emailHand := email.EmailHandler{
		From:     os.Getenv("ADDR"),
		To:       contact.GetContacts(),
		Password: os.Getenv("PASS"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
	}

	emailHand.SendCurrInfo(currData)
}
