package logger

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//aaa.bbb.ccc.ddd - - [day/month/year HH:MM:SS] "Method Path Protocol" Status -
		//100.10.0.1 - - [01/02/2023 12:13:14] "GET / HTTP/1.1" 200 -

		nl := newLoggingResponseWriter(w)
		next.ServeHTTP(nl, r)

		time := time.Now().Format("02/01/2006THH:MM:SS")
		log := fmt.Sprintf(`[ProxyLog] %s - - [%s] "%s %s %s" %d -`, r.RemoteAddr, time, r.Method, r.URL.Path, r.Proto, nl.statusCode)
		fmt.Println(log)
	})
}


func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
    return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
    lrw.statusCode = code
    lrw.ResponseWriter.WriteHeader(code)

}
