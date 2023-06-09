package proxy

import (
	"net/http"
	"strings"
)

func Proxy(handlers Proxys) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range handlers {
			if strings.Contains(r.URL.Path, k) {
				if k != "/" && r.URL.Path == k && r.URL.Path != (k + "/") {
					http.Redirect(w, r, k + "/", http.StatusFound)
					return
				}
				reverseProxy := v.get()
				reverseProxy.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
	})
}
