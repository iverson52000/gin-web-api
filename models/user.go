package models

import (
	"fmt"
	"gin-web-api/db"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

var authModel = &AuthModel{}

type UserModel struct{}

func (m UserModel) CreateUserToken(userID int64) (user User, token Token, err error) {
	if dbErr := db.GetDB().Where("id = ?", userID).First(&user).Error; dbErr != nil {
		fmt.Println("user doesn't exist")
		return
	}

	// Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(int64(userID))
	if err != nil {
		return user, token, err
	}

	saveErr := authModel.CreateAuth(userID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, nil
}
