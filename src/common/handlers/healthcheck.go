package handlers

import (
	"io"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request, version string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, version)
}
