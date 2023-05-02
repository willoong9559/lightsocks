package main

import (
	"log"

	"github.com/willoong9559/lightsocks/client"
)

func main() {
	Lsclient, err := client.NewLsClient()
	if err != nil {
		log.Panic(err)
	}
	log.Fatal(Lsclient.Listen())
}
