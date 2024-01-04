package handlers

import (
	"context"
	"crypto/tls"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	err := godotenv.Load()
	if err != nil {
		log.Print("Print Loading .env file caused error")
	}
	var connectionString string = os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(connectionString).SetTLSConfig(&tls.Config{InsecureSkipVerify: true})
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Print("Can't Connect to Database : ", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Print("Error Pinging MongoDB : ", err)
	}
	log.Print("Connected to MongoDB Atlas")
	return &DB{
		client: client,
	}
}
