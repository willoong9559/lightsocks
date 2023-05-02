package main

import (
	"log"

	"github.com/willoong9559/lightsocks/common"
	"github.com/willoong9559/lightsocks/conf"
	"github.com/willoong9559/lightsocks/server"
)

func main() {
	conf.InitConfS()
	lsServer, err := server.NewLsServer()
	if err != nil {
		log.Panic(err)
	}
	common.PrintInfo(lsServer)
	log.Fatal(lsServer.Listen())
}
