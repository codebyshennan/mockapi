package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Represents an microservice
// Also represents the database model
// Also represents the GET data
type ServiceModel struct {
	Id          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"desc" bson:"desc"`
	SwaggerRefs []primitive.ObjectID `json:"swaggerRefs" bson:"swaggerRefs"`
	CreatedAt   primitive.DateTime   `json:"createdAt" bson:"createdAt"`
	UpdatedAt   primitive.DateTime   `json:"updatedAt" bson:"updatedAt"`
}

// Represents the data used to create ServiceModel
type ServiceCreateData struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

// Represents the data used to query for ServiceModel
type ServiceQueryData struct {
	BaseQueryData
	Id primitive.ObjectID
}

// Represents the data for used to update ServiceModel
type ServiceUpdateData struct {
	Name        string             `json:"name"`
	Description string             `json:"desc"`
	SwaggerRef  primitive.ObjectID `json:"swaggerRef"`
}

// Represents a repository used to CRUD ServiceModel
type IServiceRepo interface {
	CreateService(d *ServiceCreateData) (primitive.ObjectID, error)
	GetServices(d *ServiceQueryData) ([]ServiceModel, error)
	GetOneService(d *ServiceQueryData) (*ServiceModel, error)
	UpdateService(primitive.ObjectID, *ServiceUpdateData) error
}
