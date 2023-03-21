package mocks

import (
	"github.com/codebyshennan/mockapi/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockEndPtMock struct{}

// Creates mock endpoints tied to a mock server
func (o MockEndPtMock) CreateMockEndPts(serverId primitive.ObjectID, d *[]domain.MockEndPtCreateData) ([]string, error) {
	return nil, nil
}

// Returns a list of mock endpoints of a server
func (o MockEndPtMock) GetMockEndPts(serverId primitive.ObjectID) ([]domain.MockEndPtModel, error) {
	return nil, nil
}

// Updates mock endpoints belonging to a mock server
func (o MockEndPtMock) UpdateMockEndPts(serverId primitive.ObjectID,
	create *[]domain.MockEndPtCreateData, update *[]domain.MockEndPtUpdateData) error {
	return nil
}
