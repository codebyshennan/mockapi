package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Represents an user
// Also represents the database model
// Also represents the GET data
type UserModel struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
	IsRoot    bool               `json:"isRoot" bson:"isRoot"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}

// Represents the data used to create UserModel
type UserCreateData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	IsRoot    bool   `json:"isRoot"`
}

// Represents the data used to query for UserModel
type UserQueryData struct {
	BaseQueryData
	Id    primitive.ObjectID
	Email string
}

// Represents the data for used to update UserModel
type UserUpdateData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Represents a repository used to CRUD UserModel
type IUserRepo interface {
	CreateUser(d *UserCreateData) (string, error)
	GetUsers(d *UserQueryData) ([]UserModel, error)
	GetOneUser(d *UserQueryData) (*UserModel, error)
	GetOrInsert(d *UserCreateData) (*UserModel, error)
	UpdateUser(d *UserUpdateData) error
}
