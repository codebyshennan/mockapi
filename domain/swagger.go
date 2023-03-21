package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Represents an swagger doc
// Also represents the database model
// Also represents the GET data
type SwaggerModel struct {
	Id          primitive.ObjectID        `json:"id,omitempty" bson:"_id,omitempty"`
	SwaggerSpec map[string]any            `json:"swaggerSpec" bson:"swaggerSpec"`
	Version     string                    `json:"version" bson:"version"`
	ServiceRef  primitive.ObjectID        `json:"serviceRef" bson:"serviceRef"`
	OwnerRef    primitive.ObjectID        `json:"ownerRef" bson:"ownerRef"`
	Endpoints   []SwaggerEndPointMetadata `json:"endpts" bson:"endpts"`
	CreatedAt   primitive.DateTime        `json:"createdAt" bson:"createdAt"`
	UpdatedAt   primitive.DateTime        `json:"updatedAt" bson:"updatedAt"`
}

// Represents the data used to create SwaggerModel
type SwaggerCreateData struct {
	OwnerRef   primitive.ObjectID `json:"ownerRef"`
	Version    string             `json:"version"`
	SourceType string             `json:"sourceType"`
	Source     map[string]any     `json:"source"`
	ServiceRef primitive.ObjectID `json:"serviceRef"`
}

// Represents the data used to query for SwaggerModel
type SwaggerQueryData struct {
	BaseQueryData
	Id       primitive.ObjectID
	OwnerRef primitive.ObjectID
}

// Represents the data for used to update SwaggerModel
type SwaggerUpdateData struct {
	SourceType string `json:"sourceType"`
	Source     string `json:"source"`
}

type SwaggerEndPointMetadata struct {
	EndpointRegex string `json:"endptRegex" bson:"endptRegex"`
	Method        string `json:"method" bson:"method"`
}

// Represents a repository used to CRUD SwaggerModel
type ISwaggerRepo interface {
	CreateSwagger(d *SwaggerCreateData) (string, error)
	GetSwaggers(d *SwaggerQueryData) ([]SwaggerModel, error)
	GetOneSwagger(d *SwaggerQueryData) (*SwaggerModel, error)
	UpdateSwagger(d *SwaggerUpdateData) error
}
