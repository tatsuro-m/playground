package util

import (
	"fmt"
	"net/http"
)

func NotSupportHTTPMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte(fmt.Sprintf("Not support http method %s", r.Method)))
	if err != nil {
		fmt.Println(err)
	}
}
