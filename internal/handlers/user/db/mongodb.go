package db

import (
	"context"
	"fmt"

	"github.com/tmolyakov/go-api-xmp/internal/handlers/user"
	"github.com/tmolyakov/go-api-xmp/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user due to error: %v", err)
	}

	d.logger.Debug("convert inserted id to ojbectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("failed to convert object id to hex")
}

func (d db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectId. hex: %s", id)
	}

	filter := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		// todo 404
		return u, fmt.Errorf("failed to find user by id %s, due to error %v", id, err)
	}

	if err := result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user id %s, due to error %v", id, err)
	}

	return u, nil
}

func (d db) Update(ctx context.Context, user user.User) error {
	panic("implement me")
}

func (d db) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
