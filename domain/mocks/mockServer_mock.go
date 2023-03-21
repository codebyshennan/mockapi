package mocks

import (
	"github.com/codebyshennan/mockapi/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockServerMock struct{}

func (o MockServerMock) CreateMockServer(d *domain.MockServerCreateData) (string, error) {
	return "6316fd9870b745b768fd4d38", nil
}

func (o MockServerMock) GetMockServers(d *domain.MockServerQueryData) ([]domain.MockServerGetData, error) {
	return nil, nil
}

func (o MockServerMock) GetOneMockServer(d *domain.MockServerQueryData) (domain.MockServerGetData, error) {
	return domain.MockServerGetData{}, nil
}

func (o MockServerMock) UpdateMockServer(id primitive.ObjectID, d *domain.MockServerUpdateData) error {
	return nil
}
