package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Represents an mockServer i.e. mockProfile
// Also represents the database model
type MockServerModel struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	OwnerRef  primitive.ObjectID `json:"ownerRef" bson:"ownerRef"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}

// Represents the data used to create MockServerModel
type MockServerCreateData struct {
	Name      string                `json:"name"`
	OwnerRef  primitive.ObjectID    `json:"ownerRef"`
	Responses []MockEndPtCreateData `json:"responses"`
}

// Represents the data used to query for MockServerModel
type MockServerQueryData struct {
	BaseQueryData
	Id primitive.ObjectID
}

type MockServerUpdateData struct {
	Name      string `json:"name"`
	Responses struct {
		Create []MockEndPtCreateData `json:"create"`
		Update []MockEndPtUpdateData `json:"update"`
	} `json:"responses"`
}

type MockServerGetData struct {
	MockServerModel
	EndPts []MockEndPtModel `json:"endpts"`
}

// Represents a repository used to CRUD MockServerModel
// No updates to MockServer directly
type IMockServerRepo interface {
	CreateMockServer(d *MockServerCreateData) (string, error)
	GetMockServers(d *MockServerQueryData) ([]MockServerGetData, error)
	GetOneMockServer(d *MockServerQueryData) (MockServerGetData, error)
	UpdateMockServer(id primitive.ObjectID, d *MockServerUpdateData) error
}
