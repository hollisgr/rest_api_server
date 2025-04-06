package db

import (
	"api_server/internal/user"
	"api_server/pkg/logging"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) CreateUser(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("creating user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user due to error: %v", err)
	}
	d.logger.Debug("converting inserting id to object id")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Traceln(user)
	return "", fmt.Errorf("failed to convert objectid to hex, probably oid: %s", oid)
}
func (d *db) FindUser(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectid, hex: %s", id)
	}
	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		// TODO 404
		return u, fmt.Errorf("failed to find user by id: %s due to error: %v", id, err)
	}
	err = result.Decode(&u)
	if err != nil {
		return u, fmt.Errorf("failed to decode user from DB: %s due to error: %v", id, err)
	}
	return u, nil
}
func (d *db) DeleteUser(ctx context.Context, id string) error

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
