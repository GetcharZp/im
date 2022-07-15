package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type UserRoom struct {
	UserIdentity string `bson:"user_identity"`
	RoomIdentity string `bson:"room_identity"`
	RoomType     int    `bson:"room_type"` // 房间 类型 【1-独聊房间 2-群聊房间】
	CreatedAt    int64  `bson:"created_at"`
	UpdatedAt    int64  `bson:"updated_at"`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}

func GetUserRoomByUserIdentityRoomIdentity(userIdentity, roomIdentity string) (*UserRoom, error) {
	ur := new(UserRoom)
	err := Mongo.Collection(UserRoom{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"user_identity", userIdentity}, {"room_identity", roomIdentity}}).
		Decode(ur)
	return ur, err
}

func GetUserRoomByRoomIdentity(roomIdentity string) ([]*UserRoom, error) {
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{"room_identity", roomIdentity}})
	if err != nil {
		return nil, err
	}
	urs := make([]*UserRoom, 0)
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			return nil, err
		}
		urs = append(urs, ur)
	}
	return urs, nil
}

func JudgeUserIsFriend(userIdentity1, userIdentity2 string) bool {
	// 查询 userIdentity1 单聊房间列表
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).
		Find(context.Background(), bson.D{{"user_identity", userIdentity1}, {"room_type", 1}})
	roomIdentities := make([]string, 0)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return false
	}
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			log.Printf("Decode Error:%v\n", err)
			return false
		}
		roomIdentities = append(roomIdentities, ur.RoomIdentity)
	}
	// 获取关联 userIdentity2 单聊房间个数
	cnt, err := Mongo.Collection(UserRoom{}.CollectionName()).
		CountDocuments(context.Background(), bson.M{"user_identity": userIdentity2, "room_type": 1, "room_identity": bson.M{"$in": roomIdentities}})
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return false
	}
	if cnt > 0 {
		return true
	}

	return false
}
