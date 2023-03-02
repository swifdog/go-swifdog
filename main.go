package main

import (
	"log"
	. "swifdog.io/swifdog-sdk/swifdog"
)

func main() {
	client, _ := NewBasicClient("oylI1I4N7QDtf8WuI9UItKACpnwaI69MQZhKELeRzJfzovMmZVdvFKzrOCCr9S7j")
	account, err := client.GetAccount()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(account)

	/** List account tokens
	tokens, err := client.ListAccountTokens()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(tokens)
	*/

	/* Create new account token
	str := "test by api client"
	newToken, err := client.CreateAccountToken(&str)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(newToken)
	*/

	/* Delete existing account token
	err = client.DeleteAccountTokenById("5279f81f-b50b-46e7-a4f5-7a074eb2f1e1")
	if err != nil {
		log.Fatal(err)
	}
	*/

	prjs, err := client.ListProjects()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(prjs)

}
