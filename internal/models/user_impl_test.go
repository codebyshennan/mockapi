package models

import (
	"testing"

	"github.com/codebyshennan/mockapi/domain"
	"github.com/codebyshennan/mockapi/domain/mocks"
)

var userModel = UserModel{
	coll: mocks.CollMock{},
	log:  mocks.LoggerMock{},
}

func TestCreateUser(t *testing.T) {
	str, err := userModel.CreateUser(&domain.UserCreateData{
		FirstName: "Pei Yong",
		LastName:  "Kerk",
		Email:     "peiyong.kerk@circles.asia",
	})

	if err != nil || str == "" {
		// TODO: name this properly
		t.Errorf("Test failed")
	}
}

func TestGetUsers(t *testing.T) {
	_, err := userModel.GetUsers(&domain.UserQueryData{})

	if err != nil {
		// TODO: name this properly
		t.Errorf("Test failed")
	}
}
