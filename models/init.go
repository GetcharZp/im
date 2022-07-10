package models

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo = InitMongo()
var RDB = InitRedis()

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "admin",
	}).ApplyURI("mongodb://192.168.1.8:27017"))
	if err != nil {
		log.Println("Connection MongoDB Error:", err)
		return nil
	}
	return client.Database("im")
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
