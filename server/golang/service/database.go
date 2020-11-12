package service

import (
	"github.com/ch31212y/sakuraChatV2/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func findUserFromDB(id string, projection bson.D) (*database.User, error) {
	rs := userCol.FindOne(
		ctx,
		bson.D{{"_id", id}},
		options.FindOne().SetProjection(projection),
	)
	var user *database.User
	if rs.Decode(&user) != nil {
		return user, status.New(codes.NotFound, "user not found").Err()
	}
	return user, nil
}

func findProfileFromDB(uuid string) (*database.Profile, error) {
	rs := userCol.FindOne(
		ctx,
		bson.D{{"_id", uuid}},
		options.FindOne().SetProjection(bson.D{{"profile", 1}}),
	)
	var user *database.User
	if rs.Decode(&user) != nil {
		return nil, status.New(codes.NotFound, "user not found").Err()
	}
	return &user.Profile, nil
}
