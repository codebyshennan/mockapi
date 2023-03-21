package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Represents an store record related to a mock server
// Also represents the database model
type MockServerStore struct {
	Id            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Data          string             `json:"data" bson:"data"`
	Dest          string             `json:"dest" bson:"dest"`
	MockServerRef primitive.ObjectID `json:"mockServerRef" bson:"mockServerRef"`
}

type MockServerStoreCreateData struct {
	Data string `json:"data"`
	Dest string `json:"dest"`
}

type IMockServerStoreRepo interface {
	CreateRecord(mockServerId primitive.ObjectID, data *MockServerStoreCreateData) (string, error)
	GetRecord(mockServerId primitive.ObjectID, namespace string) ([]string, error)
}
