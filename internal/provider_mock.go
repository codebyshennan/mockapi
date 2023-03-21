package internal

import (
	"bitbucket.org/libertywireless/circles-sandbox/domain"
	"bitbucket.org/libertywireless/circles-sandbox/domain/mocks"
)

func NewMockProvider() Provider {
	p := Provider{
		UserRepo:            mocks.UserMock{},
		MockServerRepo:      mocks.MockServerMock{},
		MockEndPtRepo:       mocks.MockEndPtMock{},
		MockServerStoreRepo: mocks.MockServerStoreMock{},
		ServiceRepo:         mocks.ServiceMock{},
		SwaggerRepo:         mocks.SwaggerMock{},
		Config:              &domain.Config{},
		Logger:              mocks.LoggerMock{},
	}

	return p
}
