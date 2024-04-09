package models

type ValiadateApiRequestBody struct {
	RequestPath string `json:"requestPath"`
}

type ValiadateApiResponse struct {
	Data struct {
		IsAuth bool `json:"isAuth"`
	} `json:"data"`
	IsSuccess bool   `json:"isSuccess"`
	Error     string `json:"error"`
}
