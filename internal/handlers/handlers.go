package handlers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/suhail34/notes-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DB) CreateNotes(c *gin.Context) {
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user, ok := userClaims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error claims"})
		return
	}
	var newNotes models.Notes
	if err := c.ShouldBindJSON(&newNotes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newNotes.User_name = user["user_name"].(string)
  validationErr := validate.Struct(newNotes)
  if validationErr != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
    return
  }
	collections := db.client.Database("notes").Collection("note")
	_, err := collections.InsertOne(c, newNotes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Notes created Successfully"})
}

func (db *DB) SearchTextOnNotes(c *gin.Context) {
  queryParam := c.Query("q")
  if queryParam == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "query param not available"})
    return
  }
  indexModel := mongo.IndexModel{
    Keys: bson.D{
      {"title", "text"},
      {"description", "text"},
    },
  }
  collections := db.client.Database("notes").Collection("note")
  _, err := collections.Indexes().CreateOne(c, indexModel)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
    return
  }

  cursor, err := collections.Find(c, bson.M{
    "$text": bson.M{
      "$search":queryParam,
    },
  })
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  defer cursor.Close(c)
  var results []models.Notes
  if err := cursor.All(c, &results); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusFound, gin.H{"result": results})
}

func (db *DB) GetAllNotesHandler(c *gin.Context) {
  userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user, ok := userClaims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error claims"})
		return
	}
  username := user["user_name"].(string)
  collections := db.client.Database("notes").Collection("note")
  cursor, err := collections.Find(c, bson.D{
    {"user_name", username},
  })
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  var results []models.Notes
  if err := cursor.All(c, &results); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusFound, gin.H{"results": results})
}

func (db *DB) GetNoteHandler(c *gin.Context) {
  params := c.Param("id")
  collections := db.client.Database("notes").Collection("note")
  objectID, err := primitive.ObjectIDFromHex(params)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  var result models.Notes
  err = collections.FindOne(c, bson.M{"_id": objectID}).Decode(&result)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusFound, gin.H{"result": result})
}

func (db *DB) UpdateNoteHandler(c *gin.Context) {
  params := c.Param("id")
  objectID, err := primitive.ObjectIDFromHex(params)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  var updateFields map[string]interface{}
  if err := c.ShouldBindJSON(&updateFields); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
    return
  }

  update := bson.M{"$set": updateFields}

  filter := bson.M{"_id": objectID}
  collection := db.client.Database("notes").Collection("note")
  result, err := collection.UpdateOne(c, filter, update)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"data":updateFields, "updated": result})
}

func (db *DB) DeleteNoteHandler(c *gin.Context) {
  params := c.Param("id")
  objectID, err := primitive.ObjectIDFromHex(params)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  filter := bson.M{"_id": objectID}
  collection := db.client.Database("notes").Collection("note")
  result, err := collection.DeleteOne(c, filter)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"deleted": result})
}

func (db *DB) ShareNoteHandler(c *gin.Context) {
  noteID := c.Param("id")
  var shareRequest struct {
    UsersToShareWith []string `json:"usersToShareWith" binding"required"`
  }
  if err := c.ShouldBindJSON(&shareRequest); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
    return
  }
  objectID, err := primitive.ObjectIDFromHex(noteID)
  if err!=nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Note ID"})
    return
  }
  update := bson.M{"$addToSet": bson.M{"userToShare": bson.M{"$each":shareRequest.UsersToShareWith}}}
  filter := bson.M{"_id":objectID}
  collection := db.client.Database("notes").Collection("note")
  result, err := collection.UpdateOne(c, filter, update)
  if err!=nil{
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to share Note"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message":fmt.Sprintf("Note Shared Successfully %v %v", result, noteID)})
}
