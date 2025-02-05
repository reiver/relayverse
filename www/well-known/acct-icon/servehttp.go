package verboten

import (
	"encoding/json"
	"strings"
	"net/http"
	"net/url"

	"github.com/reiver/relayverse/srv/cache"
	"github.com/reiver/relayverse/srv/http"
	"github.com/reiver/relayverse/srv/log"
)

const path string = "/.well-known/acct-icon"

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

	var queryValues url.Values
	{
		var httpRequestURL *url.URL = request.URL
		if nil == httpRequestURL {
			const code int = http.StatusInternalServerError
			http.Error(responsewriter, http.StatusText(code), code)
			log.Error("nil http-request-url")
			return
		}

		queryValues = httpRequestURL.Query()
		if nil == queryValues {
			const code int = http.StatusBadRequest
			http.Error(responsewriter, http.StatusText(code), code)
			log.Error("empty http-request query")
			return
		}
	}

	var resource string
	{
		var resources []string
		var found bool
		resources, found = queryValues["resource"]
		if !found {
			const code int = http.StatusBadRequest
			http.Error(responsewriter, http.StatusText(code), code)
			log.Error("resource not provided in acct-icon request")
			return
		}
		if 1 != len(resources) {
			const code int = http.StatusBadRequest
			http.Error(responsewriter, http.StatusText(code), code)
			log.Error("too many resources provided in acct-icon request")
			return
		}

                resource = resources[0]
	}
	log.Debugf("resource = %q", resource)
	log.Debugf("http-method = %q", request.Method)

	switch request.Method {
	case http.MethodGet:
		serveGET(responsewriter, request, resource)
		return
	default:
		const code int = http.StatusMethodNotAllowed
		http.Error(responsewriter, http.StatusText(code), code)
		log.Errorf("method %q not allowed", request.Method)
		return
	}
}

func serveGET(responsewriter http.ResponseWriter, request *http.Request, resource string) {

	log := logsrv.Prefix("www("+path+").GET").Begin()
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
		const prefix string = "acct:"
		if !strings.HasPrefix(resource, prefix) {
			const code int = http.StatusNotFound
			http.Error(responsewriter, http.StatusText(code), code)
			log.Errorf("not an acct-uri: %q", resource)
			return
		}
	}

	var body []byte
	{
		var found bool
		body, found = cachesrv.Get(resource)
		if !found {
			const code int = http.StatusNotFound
			http.Error(responsewriter, http.StatusText(code), code)
			log.Errorf("cache MISS: %q", resource)
			return
		}
		log.Debugf("cache HIT: %q", resource)
	}
	log.Debugf("body length: %d", len(body))

	type Activity struct {
		Icon struct{
			URL string `json:"url"`
		} `json:"icon"`
	}
	var activity Activity
	{
		err := json.Unmarshal(body, &activity)
		if nil != err {
			const code int = http.StatusNotFound
			http.Error(responsewriter, http.StatusText(code), code)
			log.Errorf("problem decoding JSON: %s", err)
			return
		}
	}

	var url string = activity.Icon.URL
	if "" == url {
		log.Debugf("BODY: %s", body)
	}
		
//@TODO: validate URL
		
	log.Debugf("icon-url: %q", url)

	{
		http.Redirect(responsewriter, request, url, http.StatusTemporaryRedirect)
	}
}
