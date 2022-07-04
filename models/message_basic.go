package models

import (
	"context"
)

type MessageBasic struct {
	UserIdentity string `bson:"user_identity"`
	RoomIdentity string `bson:"room_identity"`
	Data         string `bson:"data"`
	CreatedAt    int64  `bson:"created_at"`
	UpdatedAt    int64  `bson:"updated_at"`
}

func (MessageBasic) CollectionName() string {
	return "message_basic"
}

func InsertOneMessageBasic(mb *MessageBasic) error {
	_, err := Mongo.Collection(MessageBasic{}.CollectionName()).InsertOne(context.Background(), mb)
	return err
}
