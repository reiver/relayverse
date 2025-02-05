package verboten

import (
	"io"
	"strings"
	"net/http"
	"net/url"
	"time"

	"github.com/reiver/relayverse/cfg"
	"github.com/reiver/relayverse/lib/http"
	"github.com/reiver/relayverse/srv/cache"
	"github.com/reiver/relayverse/srv/http"
	"github.com/reiver/relayverse/srv/log"
)

const path string = "/.well-known/acct-cache"

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
			log.Error("resource not provided in acct-cache request")
			return
		}
		if 1 != len(resources) {
			const code int = http.StatusBadRequest
			http.Error(responsewriter, http.StatusText(code), code)
			log.Error("too many resources provided in acct-cache request")
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
	case http.MethodPut:
		servePUT(responsewriter, request, resource)
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

	{
		responsewriter.Header().Add("Content-Type", libhttp.DetectContentType(body))

		_, err := responsewriter.Write(body)
		if nil != err {
			log.Errorf("problem writing body: %s", err)
			return
		}
	}
}

func servePUT(responsewriter http.ResponseWriter, request *http.Request, resource string) {

	log := logsrv.Prefix("www("+path+").PUT").Begin()
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
	if nil == request.Header {
		const code int = http.StatusInternalServerError
		http.Error(responsewriter, http.StatusText(code), code)
		log.Error("nil http-request-headers")
		return
	}
	if nil == request.Body {
		const code int = http.StatusInternalServerError
		http.Error(responsewriter, http.StatusText(code), code)
		log.Error("nil http-request-body")
		return
	}

	
	//@TODO: authorization check
	

	var duration time.Duration = cfg.CacheDurationMinimum()

	var until time.Time = time.Now().Add(duration)
	{
		const name string = "Expires"

		var expiresString string = request.Header.Get(name)

		expires, err := http.ParseTime(expiresString)
		if nil != err {
			log.Debugf("bad %q HTTP request header value (%q): %s", name, expiresString, err)
		} else {
			if 0 < until.Compare(expires) {
				until = expires
			}
		}
	}

	var body []byte
	{
		var err error
		body, err = io.ReadAll(request.Body)
		if nil != err {
			const code int = http.StatusInternalServerError
			http.Error(responsewriter, http.StatusText(code), code)
			log.Errorf("problem reading all of http-request-body")
			return
		}
	}
	log.Debugf("body length: %d", len(body))

	{
		cachesrv.Set(resource, body, until)
	}

	{
		libhttp.NoContent(responsewriter, request)
	}
}
