package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/codebyshennan/mockapi/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// An actual instance of mockServer model tied to the db
type MockServerModel struct {
	Collection    *domain.ICollection
	MockEndPtRepo *domain.IMockEndPtRepo
}

// Factory method to create mockServer model
func NewMockServerModel(coll domain.ICollection, mockEndPtRepo domain.IMockEndPtRepo) (m *MockServerModel) {
	m = &MockServerModel{}

	// TODO
	if coll == nil || mockEndPtRepo == nil {
		fmt.Println("Missing coll or mockEndPtRepo", coll, mockEndPtRepo)
	}

	m.Collection = &coll
	m.MockEndPtRepo = &mockEndPtRepo
	return
}

// Implements IMockServerRepo interface
func (o MockServerModel) CreateMockServer(d *domain.MockServerCreateData) (string, error) {
	now := primitive.NewDateTimeFromTime(time.Now())
	db := domain.MockServerModel{
		Name:      d.Name,
		OwnerRef:  d.OwnerRef,
		CreatedAt: now,
		UpdatedAt: now,
	}

	res, err := (*o.Collection).InsertOne(context.TODO(), db)
	if err != nil {
		return "", err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "<internal_error>", nil
	}

	if _, err = (*o.MockEndPtRepo).CreateMockEndPts(id, &d.Responses); err != nil {
		return "", err
	}

	return id.Hex(), nil
}

// Implements IMockServerRepo interface
func (o MockServerModel) GetMockServers(d *domain.MockServerQueryData) ([]domain.MockServerGetData, error) {
	findOptions := options.Find()
	if d.Limit != 0 {
		findOptions.SetLimit(d.Limit)
	}
	if d.Skip != 0 {
		findOptions.SetLimit(d.Skip)
	}

	filter := map[string]any{}
	if d.Id != primitive.NilObjectID {
		filter["_id"] = d.Id
	}

	cursor, err := (*o.Collection).Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}

	var res []domain.MockServerModel
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}

	output := make([]domain.MockServerGetData, 0)
	for _, m := range res {
		if endpts, err := (*o.MockEndPtRepo).GetMockEndPts(m.Id); err != nil {
			continue
		} else {
			output = append(output, domain.MockServerGetData{
				MockServerModel: m,
				EndPts:          endpts,
			})
		}
	}

	return output, nil
}

// Implements IMockServerRepo interface
func (o MockServerModel) GetOneMockServer(d *domain.MockServerQueryData) (domain.MockServerGetData, error) {
	if d.Id == primitive.NilObjectID {
		return domain.MockServerGetData{}, errors.New("ERR_MOCKSERVER_TODO")
	}

	filter := map[string]any{
		"_id": d.Id,
	}

	var res domain.MockServerModel
	if err := (*o.Collection).FindOne(context.TODO(), filter).Decode(&res); err != nil {
		return domain.MockServerGetData{}, err
	}

	endpts, err := (*o.MockEndPtRepo).GetMockEndPts(d.Id)
	if err != nil {
		return domain.MockServerGetData{}, err
	}

	return domain.MockServerGetData{
		MockServerModel: res,
		EndPts:          endpts,
	}, nil
}

// Implement IMockServerRepo interface
func (o MockServerModel) UpdateMockServer(id primitive.ObjectID, d *domain.MockServerUpdateData) error {
	mockServer, err := o.GetOneMockServer(&domain.MockServerQueryData{
		Id: id,
	})
	if err != nil {
		return err
	}

	filter := map[string]primitive.ObjectID{
		"_id": id,
	}
	if d.Name != "" {
		mockServer.Name = d.Name
		mockServer.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	}

	if _, err := (*o.Collection).ReplaceOne(context.TODO(), filter, mockServer.MockServerModel); err != nil {
		return err
	}

	if err = (*o.MockEndPtRepo).UpdateMockEndPts(id,
		&d.Responses.Create, &d.Responses.Update); err != nil {
		return err
	}
	return nil
}
