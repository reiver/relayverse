package verboten

import (
	"net/http"

	"github.com/reiver/relayverse/srv/http"
	"github.com/reiver/relayverse/srv/log"
)

const path string = "/"

func init() {
	err := httpsrv.Mux.HandlePath(http.HandlerFunc(serveHTTP), path)
	if nil != err {
		panic(err)
	}
}

func serveHTTP(responsewriter http.ResponseWriter, request *http.Request) {
	log := logsrv.Prefix("www("+path+")").Begin()
	defer log.End()

	if nil == responsewriter {
		log.Error("nil response-writer")
		return
	}
	if nil == request {
		const code int = http.StatusInternalServerError
		http.Error(responsewriter, http.StatusText(code), code)
		log.Error("nil http-request")
		return
	}

	{
		_, err := responsewriter.Write(webpage)
		if nil != err {
			log.Errorf("problem writing http-response to http-client: %s", err)
		}
	}
}
