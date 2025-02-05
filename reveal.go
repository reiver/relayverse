package main

import (
	"time"

	"github.com/reiver/relayverse/cfg"
	"github.com/reiver/relayverse/srv/log"
)

func reveal() {
	log := logsrv.Prefix("reveal").Begin()
	defer log.End()

        var tcpaddr string = cfg.WebServerTCPAddress()
        log.Informf("serving HTTP on TCP address: %q", tcpaddr)

        var mincache time.Duration = cfg.CacheDurationMinimum()
        log.Informf("cache minimum duration to store: %v", mincache)

        var cachegcsleepmin time.Duration = cfg.CacheGCSleepMinimum()
        log.Informf("cache garbage-collector (GC) sleep minimum: %v", cachegcsleepmin)
}
