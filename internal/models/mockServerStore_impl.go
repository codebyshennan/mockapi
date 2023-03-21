package models

import (
	"context"

	"github.com/codebyshennan/mockapi/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockServerStoreRepo struct {
	logger domain.ILogger
	coll   domain.ICollection
}

func NewMockServerStoreRepo(coll domain.ICollection, logger domain.ILogger) (m MockServerStoreRepo) {
	m = MockServerStoreRepo{
		logger: logger,
		coll:   coll,
	}
	return
}

func (o MockServerStoreRepo) CreateRecord(mockServerId primitive.ObjectID,
	data *domain.MockServerStoreCreateData) (s string, err error) {

	db := domain.MockServerStore{
		Data:          data.Data,
		Dest:          data.Dest,
		MockServerRef: mockServerId,
	}

	res, err := o.coll.InsertOne(context.TODO(), db)
	if err != nil {
		return "", err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "<internal_error>", nil
	}

	return id.Hex(), nil
}

func (o MockServerStoreRepo) GetRecord(mockServerId primitive.ObjectID, namespace string) ([]string, error) {
	cursor, err := o.coll.Find(context.TODO(), map[string]any{
		"mockServerRef": mockServerId,
		"dest":          namespace,
	}, &options.FindOptions{
		Projection: map[string]bool{
			"data": true,
		},
	})
	if err != nil {
		return nil, err
	}

	data := []struct {
		Data string `json:"data" bson:"data"`
	}{}
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	output := make([]string, 0)
	for _, m := range data {
		output = append(output, m.Data[1:len(m.Data)-1])
	}

	return output, nil

}
