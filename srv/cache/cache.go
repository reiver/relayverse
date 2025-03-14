package cachesrv

import (
	"github.com/reiver/go-reg"
	"github.com/reiver/go-tmp"

	"time"
)

var cache reg.Registry[tmp.Temporal[[]byte]]

func whenCleanUp(temporal tmp.Temporal[[]byte]) bool {
	return temporal.Optional().IsNothing()
}

func Get(name string) ([]byte, bool) {
	temporal, found := cache.Get(name)
	if !found {
		var nada []byte
		return nada, false
	}

	return temporal.Optional().Get()
}

func Set(name string, value []byte, until time.Time) ([]byte, bool) {
	newTemporal := tmp.Temporary(value, until)

	temporal, found := cache.Set(name, newTemporal)
	if !found {
		var nada []byte
		return nada, false
	}

	return temporal.Get()
}

func Unset(name string) ([]byte, bool) {
	temporal, found := cache.Unset(name)
	if !found {
		var nada []byte
		return nada, false
	}

	return temporal.Get()
}
