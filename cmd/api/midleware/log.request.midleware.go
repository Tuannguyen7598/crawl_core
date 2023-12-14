package midleware

import (
	"fmt"
	"net/http"
	"time"
)

func MiddlewareGlobal(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do stuff before the handlers
		start := time.Now()

		defer func() {
			duration := time.Since(start)
			fmt.Printf("[%s] %s %s - %v\n", r.Method, r.RemoteAddr, r.URL.Path, duration)
		}()
		h.ServeHTTP(w, r)
		// do stuff after the hadlers

	})
}
