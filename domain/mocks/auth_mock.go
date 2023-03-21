package mocks

import "github.com/codebyshennan/mockapi/domain"

type AuthRepoMock struct{}

func (o AuthRepoMock) GoogleLogin(d *domain.GoogleLoginPostData) (*domain.AuthRes, error) {
	return nil, nil
}

func (o AuthRepoMock) ValidateAndDecodeJWT(token string) (*domain.UserModel, error) {
	return nil, nil
}
