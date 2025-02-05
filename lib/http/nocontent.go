package libhttp

import (
	"net/http"
)

func NoContent(responsewriter http.ResponseWriter, request *http.Request) {
	if nil == responsewriter {
		return
	}

	responsewriter.WriteHeader(http.StatusNoContent)
}
