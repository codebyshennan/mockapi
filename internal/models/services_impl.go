package models

import (
	"context"
	"errors"
	"time"

	"github.com/codebyshennan/mockapi/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// An actual instance of service model tied to the db
type ServiceModel struct {
	coll *domain.ICollection
	log  *domain.ILogger
}

// Factory method to create service model
func NewServiceModel(coll domain.ICollection, log domain.ILogger) (m *ServiceModel) {
	if coll == nil || log == nil {
		return
	}

	m = &ServiceModel{
		coll: &coll,
		log:  &log,
	}
	return
}

// Implements IServiceRepo method
func (o ServiceModel) CreateService(d *domain.ServiceCreateData) (primitive.ObjectID, error) {
	now := primitive.NewDateTimeFromTime(time.Now())

	db := domain.ServiceModel{
		Name:        d.Name,
		Description: d.Description,
		SwaggerRefs: make([]primitive.ObjectID, 0),
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	res, err := (*o.coll).InsertOne(context.TODO(), db)
	if err != nil {
		return primitive.NilObjectID, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); !ok {
		return primitive.NilObjectID, errors.New("ERR_SERVICE_TODO")
	} else {
		return oid, err
	}
}

// Implements IServiceRepo method
func (o ServiceModel) GetServices(d *domain.ServiceQueryData) ([]domain.ServiceModel, error) {
	findOptions := options.Find()
	if d.Limit != 0 {
		findOptions.SetLimit(d.Limit)
	}
	if d.Skip != 0 {
		findOptions.SetSkip(d.Skip)
	}

	filter := map[string]any{}
	if d.Id != primitive.NilObjectID {
		filter["_id"] = d.Id
	}

	cursor, err := (*o.coll).Find(context.TODO(), filter, findOptions)
	if err != nil {
		(*o.log).Errorln("ERR_SERVICE_TODO", err.Error())
		return nil, err
	}

	res := make([]domain.ServiceModel, 0)
	if err = cursor.All(context.TODO(), &res); err != nil {
		(*o.log).Errorln("ERR_SERVICE_TODO", err.Error())
		return nil, err
	}
	return res, nil
}

// Implements IServiceRepo method
func (o ServiceModel) GetOneService(d *domain.ServiceQueryData) (*domain.ServiceModel, error) {
	if d.Id == primitive.NilObjectID {
		return nil, errors.New("ERR_SERVICE_TODO")
	}

	filter := map[string]any{
		"_id": d.Id,
	}

	var res domain.ServiceModel
	if err := (*o.coll).FindOne(context.TODO(), filter).Decode(&res); err != nil {
		(*o.log).Errorln("ERR_SERVICE_TODO", err.Error())
		return nil, err
	} else {
		return &res, nil
	}
}

// Implements IServiceRepo method
func (o ServiceModel) UpdateService(id primitive.ObjectID, d *domain.ServiceUpdateData) error {
	service, err := o.GetOneService(&domain.ServiceQueryData{
		Id: id,
	})
	if err != nil {
		return err
	}

	filter := map[string]primitive.ObjectID{
		"_id": id,
	}

	if d.Name != "" {
		service.Name = d.Name
		service.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	}
	if d.Description != "" {
		service.Description = d.Description
		service.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	}
	if d.SwaggerRef != primitive.NilObjectID {
		service.SwaggerRefs = append(service.SwaggerRefs, d.SwaggerRef)
		service.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	}

	if _, err := (*o.coll).ReplaceOne(context.TODO(), filter, service); err != nil {
		return err
	}

	return nil
}
