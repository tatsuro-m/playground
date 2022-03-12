package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func LoggingReq(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "[%s][%s]: %s\n", r.Method, r.URL, time.Now().Format(time.RFC3339))
		next.ServeHTTP(w, r)
	}
}
