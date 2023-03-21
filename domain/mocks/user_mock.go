package mocks

import "github.com/codebyshennan/mockapi/domain"

type UserMock struct{}

func (o UserMock) CreateUser(d *domain.UserCreateData) (string, error) {
	return "6316fd387f943268924e0642", nil
}

func (o UserMock) GetUsers(d *domain.UserQueryData) ([]domain.UserModel, error) {
	return nil, nil
}

func (o UserMock) GetOneUser(d *domain.UserQueryData) (*domain.UserModel, error) {
	return &domain.UserModel{}, nil
}

func (o UserMock) GetOrInsert(d *domain.UserCreateData) (*domain.UserModel, error) {
	return nil, nil
}

func (o UserMock) UpdateUser(d *domain.UserUpdateData) error {
	return nil
}
