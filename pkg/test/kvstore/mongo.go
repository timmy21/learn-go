package kvstore

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ Backend = (*MongoBackend)(nil)

const tableName = "kvstore"

type MongoBackend struct {
	database *mongo.Database
}

func NewMongoBackend(db *mongo.Database) *MongoBackend {
	return &MongoBackend{database: db}
}

func (m *MongoBackend) Set(ctx context.Context, key string, value []byte) error {
	_, err := m.database.Collection(tableName).UpdateByID(
		ctx,
		key,
		bson.D{{
			Key: "$set", Value: bson.D{{
				Key: "value", Value: value}}}},
		options.Update().SetUpsert(true),
	)
	return errors.WithStack(err)
}

func (m *MongoBackend) Get(ctx context.Context, key string) ([]byte, error) {
	result := m.database.Collection(tableName).FindOne(
		ctx,
		bson.D{{Key: "_id", Value: key}},
		options.FindOne().SetProjection(bson.D{{Key: "value", Value: 1}}),
	)
	switch err := result.Err(); {
	case errors.Is(mongo.ErrNoDocuments, err):
		return nil, errors.WithStack(&NotFoundError{key: key})
	case err != nil:
		return nil, errors.WithStack(err)
	default:
		var doc struct {
			Value []byte `bson:"value"`
		}
		err := result.Decode(&doc)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return doc.Value, nil
	}
}
