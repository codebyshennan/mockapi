package mocks

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollMock struct{}

func (o CollMock) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return &mongo.InsertOneResult{
		InsertedID: primitive.NewObjectID(),
	}, nil
}

func (o CollMock) UpdateOne(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return nil, nil
}

func (o CollMock) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	curs, err := mongo.NewCursorFromDocuments([]any{}, nil, nil)
	return curs, nil
}

func (o CollMock) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {
	return nil
}

func (o CollMock) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{},
	opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return nil, nil
}

func (o CollMock) DeleteOne(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return nil, nil
}
