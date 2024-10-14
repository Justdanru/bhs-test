package models

type UserCredentials struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}
