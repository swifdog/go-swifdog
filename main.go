package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Account struct {
	Id               string `json:"id"`
	Email            string `json:"email"`
	CreationDateTime string `json:"creationDateTime"`
}

func main() {
	client, _ := NewBearerTokenClient("oylI1I4N7QDtf8WuI9UItKACpnwaI69MQZhKELeRzJfzovMmZVdvFKzrOCCr9S7j")
	account, err := client.GetAccount()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(account)
}
