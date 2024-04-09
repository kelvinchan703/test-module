package middleware

import (
	commonConfig "ctint-conv/src/common/config"
	commonUtils "ctint-conv/src/common/utils"

	"net/http"
)

// CheckHeader is a middleware function to check for necessary headers
func ValidateHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, header := range commonConfig.GlobalConfig.Auth.Headers {
			value := r.Header.Get(header.Name)
			if value == "" {
				commonUtils.HandleErrorWithStatus(w, commonUtils.NewError("Header "+header.Name+" is required."), http.StatusBadRequest)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func SetGlobalHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the desired header
		w.Header().Set("Content-Type", "application/json")

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
