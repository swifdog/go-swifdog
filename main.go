package main

import (
	"log"
	. "swifdog.io/swifdog-sdk/swifdog"
)

func main() {
	client, _ := NewBearerTokenClient("bF22pK1Kal7aqFgpMvYX1t22ILS5JVlHVvy1SDzihRdPTkBK1gnwSeiuXCEdHAaP")
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

	for _, prj := range prjs {
		err = client.DeleteProjectById(prj.ID)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println(prjs)

	newPrj, err := client.CreateProject(&CreateOrPatchProjectRequest{

	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(newPrj)
	newPrj, err = client.PatchProject(newPrj.ID, &CreateOrPatchProjectRequest{
		Name: "test-project",
		Description: "ich mag golang!",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(newPrj)

	newVolume, err := client.CreatePersistentVolume(newPrj.ID, &CreateOrPatchPersistentVolumeRequest{
		Name: "demo-volume",
		Capacity: "1G",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(newVolume)

	volumes, err := client.ListPersistentVolume(newPrj.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(volumes)
}
