package models

import (
	"context"

	"github.com/codebyshennan/mockapi/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// An actual instance of mockEndPt model tied to the db
type MockEndPtModel struct {
	coll *domain.ICollection
}

// Factory method to create mockEndPt model
func NewMockEndPtModel(coll domain.ICollection) (m *MockEndPtModel) {
	m = &MockEndPtModel{
		coll: &coll,
	}

	return
}

// Implements IMockEndPtRepo interface
func (o MockEndPtModel) CreateMockEndPts(server primitive.ObjectID,
	d *[]domain.MockEndPtCreateData) ([]string, error) {

	ids := make([]string, 0)
	for _, c := range *d {
		db := domain.MockEndPtModel{
			MockServerRef: server,
			EndpointRegex: c.EndpointRegex,
			Method:        c.Method,
			ResponseCode:  c.ResponseCode,
		}

		if c.ResponseBody != nil {
			db.ResponseBody = c.ResponseBody
		}
		if c.Timeout != nil {
			db.Timeout = c.Timeout
		}
		if c.Writes != nil {
			db.Writes = &domain.MockEndPtWriteField{
				Dest: string("default"),
				Data: c.Writes.Data,
			}
		}

		res, err := (*o.coll).InsertOne(context.TODO(), db)
		if err != nil {
			return nil, err
		}

		if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
			ids = append(ids, oid.Hex())
		}
	}

	return ids, nil
}

// Implements IMockEndPtRepo interface
func (o MockEndPtModel) GetMockEndPts(serverId primitive.ObjectID) ([]domain.MockEndPtModel, error) {
	cursor, err := (*o.coll).Find(context.TODO(), map[string]primitive.ObjectID{
		"mockServerRef": serverId,
	})
	if err != nil {
		return nil, err
	}

	res := make([]domain.MockEndPtModel, 0)
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

// Implements IMockEndPtRepo interface
func (o MockEndPtModel) UpdateMockEndPts(serverId primitive.ObjectID,
	create *[]domain.MockEndPtCreateData, update *[]domain.MockEndPtUpdateData) error {

	endpts, err := o.GetMockEndPts(serverId)
	if err != nil {
		return err
	}

	// update first
	for _, u := range *update {
		oid, err := primitive.ObjectIDFromHex(u.Id)
		if err != nil {
			return err
		}

		for i := range endpts {
			if endpts[i].Id != oid {
				continue
			}

			endpts[i].Method = u.Method
			endpts[i].EndpointRegex = u.EndpointRegex

			if u.ResponseBody != nil {
				endpts[i].ResponseBody = u.ResponseBody
			}
			if u.ResponseCode != 0 {
				endpts[i].ResponseCode = u.ResponseCode
			}
			if u.Timeout != nil {
				endpts[i].Timeout = u.Timeout
			}
			if u.Writes != nil {
				endpts[i].Writes = &domain.MockEndPtWriteField{
					Data: u.Writes.Data,
					Dest: u.Writes.Dest,
				}
			}

			if _, err := (*o.coll).ReplaceOne(context.TODO(), map[string]primitive.ObjectID{
				"_id": oid,
			}, endpts[i]); err != nil {
				return err
			}
		}
	}

	// TODO: corner cases of overlapping endpoint regex
	if _, err := o.CreateMockEndPts(serverId, create); err != nil {
		return err
	}

	return nil
}
