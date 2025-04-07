package db

import (
	"context"
	"fmt"
	"rest_api_server/internal/user"
	"rest_api_server/pkg/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

func (d *db) CreateUser(ctx context.Context, user user.User) (string, error) {

	d.logger.Debug("creating user")

	count, countErr := d.collection.CountDocuments(ctx, bson.M{})
	if countErr != nil {
		d.logger.Infoln("userlist is empty")
	}
	user.UID = count + 1
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

// func (d *db) FindUser(ctx context.Context, id string) (u user.User, err error) {
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return u, fmt.Errorf("failed to convert hex to objectid, hex: %s", id)
// 	}
// 	filter := bson.M{"_id": oid}

// 	result := d.collection.FindOne(ctx, filter)
// 	if result.Err() != nil {
// 		// TODO 404
// 		return u, fmt.Errorf("failed to find user by id: %s due to error: %v", id, err)
// 	}
// 	err = result.Decode(&u)
// 	if err != nil {
// 		return u, fmt.Errorf("failed to decode user from DB: %s due to error: %v", id, err)
// 	}
// 	return u, nil
// }

func (d *db) FindUser(ctx context.Context, id int64) (u user.User, err error) {

	filter := bson.M{"uid": id}

	d.logger.Infoln("searching for: ", id)

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		// TODO 404
		return u, fmt.Errorf("failed to find user by id: %d due to error: %v", id, err)
	}
	err = result.Decode(&u)
	if err != nil {
		return u, fmt.Errorf("failed to decode user from DB: %d due to error: %v", id, err)
	}
	return u, nil
}

func (d *db) FindAllUsers(ctx context.Context) (u []user.User, err error) {

	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		// TODO 404
		return u, fmt.Errorf("failed to find users due to error: %v", err)
	}

	err = cursor.All(ctx, &u)
	if err != nil {
		return u, fmt.Errorf("failed to read all docs from cursor, error %v", err)
	}

	return u, nil
}

func (d *db) DeleteUser(ctx context.Context, id int64) error {
	filter := bson.M{"uid": id}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query, error: %v", err)
	}
	if result.DeletedCount == 0 {
		// TODO ErrEntityNotFound
		return fmt.Errorf("not found")
	}
	d.logger.Tracef("deleted %d documents", result.DeletedCount)
	return nil
}
