package proxy

import (
	"net/http"
	"strings"
)

func Proxy(handlers Proxys) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bo bool = false
		for k, v := range handlers {
			if strings.Contains(r.URL.Path, k) {
				bo = true
				if r.URL.Path == k && r.URL.Path != (k + "/") {
					http.Redirect(w, r, k + "/", http.StatusFound)
				}
				v.ServeHTTP(w, r)
			}
		}

		if !bo {
			w.WriteHeader(http.StatusNotFound)
		}
	})
}
