package env

import (
	"os"
	"time"

	"github.com/reiver/go-erorr"
)

var CacheMin time.Duration = cacheMin()

func cacheMin() time.Duration {
	const name string = "CACHE_MIN"

	cacheMin := os.Getenv(name)
	if "" == cacheMin {
		cacheMin = "1m" // 1 minute
	}

	duration, err := time.ParseDuration(cacheMin)
	if nil != err {
		var msg error = erorr.Errorf("problem parsing (duration) value (%q) of operating-system environment-variable named %q: %s", cacheMin, name, err)
		panic(msg)
	}

	return duration
}
