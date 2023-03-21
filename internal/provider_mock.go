package internal

import (
	"github.com/codebyshennan/mockapi/domain"
	"github.com/codebyshennan/mockapi/domain/mocks"
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
