package models

type UserCredentials struct {
	Login    string `json:"login,required"`
	Password string `json:"password,required"`
}
