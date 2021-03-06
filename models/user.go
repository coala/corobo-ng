package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ProviderId int64  `json:"providerID"`
}
