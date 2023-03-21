package mocks

import "github.com/codebyshennan/mockapi/domain"

type SwaggerMock struct{}

func (o SwaggerMock) CreateSwagger(d *domain.SwaggerCreateData) (string, error) {
	return "6316fd12167edd8b14a2752d", nil
}
func (o SwaggerMock) GetSwaggers(d *domain.SwaggerQueryData) ([]domain.SwaggerModel, error) {
	return []domain.SwaggerModel{}, nil
}

func (o SwaggerMock) GetOneSwagger(d *domain.SwaggerQueryData) (*domain.SwaggerModel, error) {
	return &domain.SwaggerModel{}, nil
}

func (o SwaggerMock) UpdateSwagger(d *domain.SwaggerUpdateData) error {
	return nil
}
