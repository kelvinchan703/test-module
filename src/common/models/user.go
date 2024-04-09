package models

type UserData struct {
	GrantType string `json:"grant_type"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type UserApiResponse struct {
	UserData  UserData `json:"data"`
	IsSuccess bool     `json:"isSuccess"`
}
