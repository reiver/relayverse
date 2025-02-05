package env

import (
	"os"
	"time"

	"github.com/reiver/go-erorr"
)
var CacheGCSleep time.Duration = cacheGCSleep()

func cacheGCSleep() time.Duration {
	const name string = "CACHE_GC_SLEEP"

	cacheGCSleep := os.Getenv(name)
	if "" == cacheGCSleep {
		cacheGCSleep = "1223s" // 1223 seconds == 20 minutes 23 seconds
	}

	duration, err := time.ParseDuration(cacheGCSleep)
	if nil != err {
		var msg error = erorr.Errorf("problem parsing (duration) value (%v) of operating-system environment-variable named %q: %s", cacheGCSleep, name, err)
		panic(msg)
	}

	return duration
}
