package cfg

import (
	"github.com/reiver/relayverse/flg"

	"time"
)

func CacheDurationMinimum() time.Duration {
	return flg.CacheDurationMinimum
}
