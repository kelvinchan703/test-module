package middleware

import (
	commonUtils "ctint-conv/src/common/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// GlobalErrorHandler is a middleware to handle panics and errors globally.
func GlobalErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				// Handle panic and return JSON error response
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				err := commonUtils.NewAppError(fmt.Sprintf("%v", r))
				json.NewEncoder(w).Encode(map[string]interface{}{
					"data":      nil,
					"isSuccess": false,
					"error":     err.Error(),
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}
