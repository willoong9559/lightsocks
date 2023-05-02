package main

import (
	"log"

	"github.com/willoong9559/lightsocks/client"
	"github.com/willoong9559/lightsocks/common"
	"github.com/willoong9559/lightsocks/conf"
)

func main() {
	conf.InitConfC()
	lsClient, err := client.NewLsClient()
	if err != nil {
		log.Panic(err)
	}
	common.PrintInfo(lsClient)
	log.Fatal(lsClient.Listen())
}
