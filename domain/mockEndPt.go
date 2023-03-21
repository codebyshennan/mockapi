package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Represents an mockEndPt
// Also represents the database model
// Also represents the GET data
type MockEndPtModel struct {
	Id            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	MockServerRef primitive.ObjectID `json:"mockServerRef" bson:"mockServerRef"`
	Method        string             `json:"method" bson:"method"`
	EndpointRegex string             `json:"endptRegex" bson:"endpointRegex"`
	ResponseCode  int                `json:"resCode" bson:"resCode"`

	// Can be null
	ResponseBody *string `json:"resBody" bson:"resBody"`

	// Can be null
	Timeout *int `json:"timeout" bson:"timeout"`

	// Can be null
	// If its not null, it will have 2 keys
	Writes *MockEndPtWriteField `json:"writes" bson:"writes"`
}

// Represents the data used to create MockEndPtModel
type MockEndPtCreateData struct {
	Method        string `json:"method"`
	EndpointRegex string `json:"endptRegex"`
	ResponseCode  int    `json:"resCode"`

	// Can be null
	ResponseBody *string `json:"resBody"`

	// Can be null
	Timeout *int `json:"timeout"`

	// Can be null
	// If its not null, it will have 2 keys
	Writes *MockEndPtWriteField `json:"writes"`
}

// Represents the data for used to update MockEndPtModel
type MockEndPtUpdateData struct {
	Id            string `json:"id"`
	Method        string `json:"method"`
	EndpointRegex string `json:"endptRegex"`
	ResponseCode  int    `json:"resCode"`

	// Can be null
	ResponseBody *string `json:"resBody"`

	// Can be null
	Timeout *int `json:"timeout"`

	// Can be null
	// If its not null, it will have 2 keys
	Writes *MockEndPtWriteField `json:"writes"`
}

type MockEndPtWriteField struct {
	Dest string `json:"dest" bson:"dest"`
	Data string `json:"data" bson:"data"`
}

// Represents a repository used to CRUD MockEndPtModel
type IMockEndPtRepo interface {
	// Creates mock endpoints tied to a mock server
	CreateMockEndPts(serverId primitive.ObjectID, d *[]MockEndPtCreateData) ([]string, error)

	// Returns a list of mock endpoints of a server
	GetMockEndPts(serverId primitive.ObjectID) ([]MockEndPtModel, error)

	// Updates mock endpoints belonging to a mock server
	UpdateMockEndPts(serverId primitive.ObjectID,
		create *[]MockEndPtCreateData, update *[]MockEndPtUpdateData) error
}
