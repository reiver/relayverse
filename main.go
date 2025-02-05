package main

import (
	"github.com/reiver/relayverse/srv/log"
)

func main() {
	log := logsrv.Prefix("main").Begin()
	defer log.End()

	log.Inform("relayverse 🐙")
	shout()

	reveal()

	log.Inform("Here we go…")
	webserve()
}
