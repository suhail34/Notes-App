package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	User_name string             `json:"user_name" validate:"required,min=3,max=100"`
	Email     string             `json:"email" validate:"email,required"`
	Password  string             `json:"password", validate:"required,min=4"`
}

type Notes struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title" validate:"required,max=50"`
	Description string             `json:"description" validate:"required"`
	User_name   string             `json:"user_name"`
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
