package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/suhail34/notes-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

var validate = validator.New()

func (db *DB) SignUpHandler(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(newUser)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	if newUser.User_name == "" || newUser.Password == "" || newUser.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username, Email and Password are required"})
		return
	}
	if err := newUser.SetPassword(newUser.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	collection := db.client.Database("notes").Collection("user")
	exits, err := userNameExists(c, collection, newUser.User_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if exits {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}
	_, err = collection.InsertOne(c, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create User"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User Created Successfully"})
}

func (db *DB) LoginHandler(c *gin.Context) {
	var credentials struct {
		Username string `json:"user_name" validate:"required,min=3,max=100" binding:"required"`
		Password string `json:"password" validate:"required,min=4" binding:"required"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(credentials)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	collections := db.client.Database("notes").Collection("user")
	exists, err := userNameExists(c, collections, credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	var user models.User
	filter := bson.D{{"user_name", credentials.Username}}
	err = collections.FindOne(c, filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if err := checkPassword(user.Password, credentials.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, err := generateToken(credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
