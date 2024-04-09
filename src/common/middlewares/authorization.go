package middleware

import (
	"bytes"
	commonConfig "ctint-conv/src/common/config"
	commonModels "ctint-conv/src/common/models"
	commonUtils "ctint-conv/src/common/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ValidateCdssToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Print current route URI
		routePattern := r.URL.Path
		fmt.Printf("Current route: %s\n", routePattern)

		// Define your JSON data
		reqBody := commonModels.ValiadateApiRequestBody{
			RequestPath: routePattern,
		}

		// Marshal the struct into JSON
		jsonData, err := json.Marshal(reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a request with the JSON body
		req, err := http.NewRequest(http.MethodPost, commonConfig.GlobalConfig.Services.Auth.Host+commonConfig.GlobalConfig.Services.Auth.Basepath+"/validate", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set Content-Type header to application/json
		req.Header.Set("traceId", r.Header.Get("traceId"))
		req.Header.Set("tenant", r.Header.Get("tenant"))
		req.Header.Set("sourceId", r.Header.Get("sourceId"))
		req.Header.Set("previousId", r.Header.Get("previousId"))
		req.Header.Set("cdss_authorization", r.Header.Get("cdss_authorization"))

		// Send the request to the endpoint
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			http.Error(w, fmt.Sprintf("Unexpected status code: %d", resp.StatusCode), resp.StatusCode)
			return
		}

		// Read the response body into a bytes.Buffer
		var bodyBuffer bytes.Buffer
		_, err = io.Copy(&bodyBuffer, resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// Convert the bytes.Buffer to a string and print it
		bodyString := bodyBuffer.String()
		var response commonModels.ValiadateApiResponse
		myerr := json.Unmarshal([]byte(bodyString), &response)
		if myerr != nil {
			fmt.Println("Error:", myerr)
			return
		}

		fmt.Println("ValiadateApiResponse:", response.Data.IsAuth)
		if response.Data.IsAuth {
			next.ServeHTTP(w, r)
		} else {
			commonUtils.HandleErrorWithStatus(w, commonUtils.NewError("Forbidden Access"), http.StatusForbidden)
		}
	})
}
