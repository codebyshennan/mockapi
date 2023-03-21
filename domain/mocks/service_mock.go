package mocks

import (
	"github.com/codebyshennan/mockapi/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceMock struct{}

func (o ServiceMock) CreateService(d *domain.ServiceCreateData) (primitive.ObjectID, error) {
	return primitive.NewObjectID(), nil
}

func (o ServiceMock) GetServices(d *domain.ServiceQueryData) ([]domain.ServiceModel, error) {
	return nil, nil
}

func (o ServiceMock) GetOneService(d *domain.ServiceQueryData) (*domain.ServiceModel, error) {
	return nil, nil
}

func (o ServiceMock) UpdateService(primitive.ObjectID, *domain.ServiceUpdateData) error {
	return nil
}
