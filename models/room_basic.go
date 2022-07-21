package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type RoomBasic struct {
	Identity     string `bson:"identity"`
	Number       string `bson:"number"`
	Name         string `bson:"name"`
	Info         string `bson:"info"`
	UserIdentity string `bson:"user_identity"`
	CreatedAt    int64  `bson:"created_at"`
	UpdatedAt    int64  `bson:"updated_at"`
}

func (RoomBasic) CollectionName() string {
	return "room_basic"
}

func InsertOneRoomBasic(rb *RoomBasic) error {
	_, err := Mongo.Collection(RoomBasic{}.CollectionName()).InsertOne(context.Background(), rb)
	return err
}

func DeleteRoomBasic(roomIdentity string) error {
	_, err := Mongo.Collection(RoomBasic{}.CollectionName()).
		DeleteOne(context.Background(), bson.M{"identity": roomIdentity})
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return err
	}
	return nil
}
