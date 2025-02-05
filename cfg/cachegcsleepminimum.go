package cfg

import (
	"github.com/reiver/relayverse/flg"

	"time"
)

func CacheGCSleepMinimum() time.Duration {
	return flg.CacheGCSleepMinimum
}
