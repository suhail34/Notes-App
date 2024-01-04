package handlers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func userNameExists(c *gin.Context, collection *mongo.Collection, username string) (bool, error) {
	filter := bson.D{{"user_name", username}}
	err := collection.FindOne(c, filter).Err()
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_name": username,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	})
	secretKey := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func checkPassword(password string, cpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(cpassword))
}
