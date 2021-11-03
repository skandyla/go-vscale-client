// example usage of our vscale sdk
//https://developers.vscale.io/documentation/api/v1/

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skandyla/go-client-vscale/vscale"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatalf("Error: %s absent", token)
	}

	vscaleClient := vscale.NewClient(token)

	//---- list locations
	locations, err := vscaleClient.Region.List()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Vscale locations:", locations)
	for _, location := range locations {
		fmt.Println("Locations:", location.ID)
	}

	//---- list scalets
	scalets, err := vscaleClient.Scalet.List()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Vscale scalets:", scalets)
	for _, scalet := range scalets {
		fmt.Println("Scalet:", scalet.Name, scalet.Created, scalet.Location, scalet.PublicAddress.Address)
	}

	//---- create scalet
	//-d '{"make_from":"ubuntu_14.04_64_002_master","rplan":"medium","do_start":true,"name":"New-Test","keys":[16],"location":"spb0"}'
	createParams := vscale.ScaletCreateRequest{
		Name:     "test",
		MakeFrom: "debian_10_64_001_master",
		Rplan:    "small",
		DoStart:  true,
		Keys:     []int64{53551}, //curl 'https://api.vscale.io/v1/sshkeys' -H "X-Token: $TOKEN"
		Location: "spb0",
	}
	scalet, err := vscaleClient.Scalet.Create(&createParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created scalet:", scalet)
	fmt.Println("Scalet:", scalet.Name, scalet.Created, scalet.PublicAddress.Address)

}
