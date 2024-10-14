package models

import "github.com/Justdanru/bhs-test/internal/models"

type User struct {
	Id       uint64 `json:"id,required"`
	Username string `json:"username,required"`
}

func NewUserFromModel(user *models.User) *User {
	return &User{
		Id:       user.Id(),
		Username: user.Username(),
	}
}
