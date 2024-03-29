package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id          int64  `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    []byte `json:"password"`
	Role        string `json:"role"`
	CreatedAt   int64  `json:"createdAt" gorm:"autoUpdateTime:milli"`
	ModifiedAt  int64  `json:"modifiedAt" gorm:"autoUpdateTime:milli"`
}

func (user *User) HashPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
