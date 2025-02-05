package flg

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/reiver/relayverse/env"
)

var (
	help bool
)

var (
	CacheDurationMinimum time.Duration
	CacheGCSleepMinimum time.Duration
	WebServerTCPAddress string
)

func init() {
	var defaultCacheDurationMinimum time.Duration = env.CacheMin
	var defaultCacheGCSleepMinimum  time.Duration = env.CacheGCSleep
	var defaultWebServerTCPAddress  string        = fmt.Sprintf(":%s", env.TcpPort)

	flag.BoolVar(&help, "help", false, "prints this (help) message")
	flag.DurationVar(&CacheDurationMinimum, "cache-min",      defaultCacheDurationMinimum, "minimum cache duration (ex: --cache-min=10m30s)")
	flag.DurationVar(&CacheGCSleepMinimum,  "cache-gc-sleep", defaultCacheGCSleepMinimum,  "minimum cache garbage-collector sleep (ex: --cache-gc-sleep=2m12s)")
	flag.StringVar(&WebServerTCPAddress, "tcp-addr", defaultWebServerTCPAddress, "web-server TCP address (ex: --tcp-addr=127.0.0.1:8765)")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
		return
	}
}
