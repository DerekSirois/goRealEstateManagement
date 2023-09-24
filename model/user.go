package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password []byte
	Active   bool
}

type UserJson struct {
	Username string
	Password string
}

func (u *UserJson) MapToUser() *User {
	hashed, _ := HashPassword(u.Password)
	return &User{
		Username: u.Username,
		Password: hashed,
		Active:   true,
	}
}

func (u *User) Create(db *gorm.DB) error {
	result := db.Create(u)
	return result.Error
}

func (u *User) GetByUsername(username string, db *gorm.DB) error {
	result := db.Find(u, "username = ?", username)
	return result.Error
}

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
