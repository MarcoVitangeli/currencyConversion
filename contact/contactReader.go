package contact

import (
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func GetContacts() []string {
	data, err := os.ReadFile(os.Getenv("CONTACT_LIST_PATH"))
	check(err)
	return strings.Split(string(data), "\n")
}
