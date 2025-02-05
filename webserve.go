package main

import (
	"net/http"

	"github.com/reiver/relayverse/cfg"
	"github.com/reiver/relayverse/srv/http"
	"github.com/reiver/relayverse/srv/log"
	_ "github.com/reiver/relayverse/www" // This import enables all the HTTP handlers.
)

func webserve() {
	log := logsrv.Prefix("webserve").Begin()
	defer log.End()

	var tcpaddr string = cfg.WebServerTCPAddress()

	err := http.ListenAndServe(tcpaddr, &httpsrv.Mux)
	if nil != err {
		log.Errorf("ERROR: problem with serving HTTP on TCP address %q: %s", tcpaddr, err)
		panic(err)
	}
}
