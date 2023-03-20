package main

import (
	"log"
	"swifdog.io/swifdog-sdk/swifdog"
)

func main() {
	client, _ := swifdog.NewBasicClient("max@mustermann.de", "test")
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

	// cleanup existing projects
	/* prjs, err := client.ListProjects()
	if err != nil {
		log.Fatal(err)
	}

	for _, prj := range prjs {
		err = client.DeleteProjectById(prj.ID)
		if err != nil {
			log.Fatal(err)
		}
	}

	// create project
	newPrj, err := client.CreateProject(&swifdog.CreateOrPatchProjectRequest{
		Name: "ich-mag-kekse",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(newPrj)

	// patch project attributes
	newPrj, err = client.PatchProject(newPrj.ID, &swifdog.CreateOrPatchProjectRequest{
		Name:        "test-project",
		Description: "ich mag golang!",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(newPrj)

	// create pv
	newVolume, err := client.CreatePersistentVolume(newPrj.ID, &swifdog.CreateOrPatchPersistentVolumeRequest{
		Name:     "demo-volume",
		Capacity: "test",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(newVolume)

	// print a list of pvs
	volumes, err := client.ListPersistentVolume(newPrj.ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(volumes)

	packet, err := client.CreatePacket(newPrj.ID, &swifdog.CreateOrPatchPacketRequest{
		Name:  "nginx",
		Image: "nginx:latest",
		EnvironmentVariables: []swifdog.EnvironmentVariable{
			{
				Key:   "DATABASE_PASSWORD",
				Value: "IchBinSicher123!",
			},
		},
		VolumeMounts: []swifdog.PersistentVolumeMount{
			{
				VolumeName: newVolume.Name,
				MountPath:  "/var/www/html",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(packet) */

	newPrj, err := client.GetProject("0411676a-fe02-4961-bb81-5e994cdf9c4c")

	if err != nil {
		log.Fatal(err)
	}
	log.Println(newPrj)

	packet, err := client.GetPacket(newPrj.ID, "4c0dfdbb-02ef-41f7-baef-3d3e2f9213c5")

	if err != nil {
		log.Fatal(err)
	}
	log.Println(packet)

	ingressRule, err := client.CreateIngressRule(newPrj.ID, &swifdog.IngressRule{
		Hostname: "example.com",
		PathRules: []swifdog.IngressRulePath{
			{
				Path:          "/",
				PacketId:      packet.ID,
				ContainerPort: 80,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(ingressRule)
}
