package models

import (
	"errors"
	"task-5-pbi-btpns-Berlian/database"
	"task-5-pbi-btpns-Berlian/utils/token"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null" json:"username"`
	Email    string `gorm:"size:255;unique;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
}

func LoginCheck(email string, password string) (string, error) {
	var err error
	user := User{}
	err = database.DB.Model(User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func(u *User) SaveUser() (*User, error) {
	var err error
	err = database.DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func GetUserByID(uid uint) (User, error) {
	var user User
	if err := database.DB.First(&user, uid).Error; err != nil {
		return user, errors.New("User not found!")
	}

	return user, nil
}
