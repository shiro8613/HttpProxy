package proxy

import (
	"net/http"
	"net/http/httputil"
	"strings"
)

func (pw ProxyWare) get() httputil.ReverseProxy{
	return httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL.Scheme = pw.schem
			request.URL.Host = pw.host
			request.URL.Path = strings.Replace(request.URL.Path, pw.path, "", -1)
		},
	}
}