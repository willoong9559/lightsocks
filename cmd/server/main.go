package main

import (
	"log"

	"github.com/willoong9559/lightsocks/server"
)

func main() {
	LsServer, err := server.NewLsServer()
	if err != nil {
		log.Panic(err)
	}
	log.Fatal(LsServer.Listen())
}
