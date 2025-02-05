package cachesrv

import (
	"math/rand"
	"time"

	"github.com/reiver/go-tmp"

	"github.com/reiver/relayverse/cfg"
	"github.com/reiver/relayverse/srv/log"
)



func init() {
	go garbageCollector()
}

func garbageCollector() {

	log := logsrv.Prefix("cachesrv-garbage-collector").Begin()
	defer log.End()

	for {
		var sleep time.Duration = cfg.CacheGCSleepMinimum() + (time.Second * time.Duration(rand.Intn(59)))
		log.Debugf("sleeping for %v", sleep)
		time.Sleep(sleep)
		log.Debug("awoken")

		log.Debug("clean-up BEGUN")
		cleanUp()
		log.Debug("clean-up ENDED")
	}

}

func cleanUp() {

	log := logsrv.Prefix("cachesrv-clean-up").Begin()
	defer log.End()

	var names []string
	{
		cache.For(func(name string, _ tmp.Temporal[[]byte]){
			names = append(names, name)
		})
	}
	log.Debugf("number of names: %d", len(names))
	if length := len(names); 0 < length && length < 14 {
		log.Debugf("names: %+v", names)
	}

	for _, name := range names {
		var sleep time.Duration = time.Nanosecond * time.Duration(rand.Intn(1024))
		time.Sleep(sleep)

		_, _, wasUnset := cache.UnsetWhen(name, whenCleanUp)
		if wasUnset {
			log.Debugf("unset %q", name)
		}
	}
}
