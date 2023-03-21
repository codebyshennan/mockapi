package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/codebyshennan/mockapi/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// An actual instance of swagger model tied to the db
type SwaggerModel struct {
	coll        *domain.ICollection
	log         *domain.ILogger
	serviceRepo *domain.IServiceRepo
}

// Factory method to create swagger model
func NewSwaggerModel(coll domain.ICollection, log *domain.ILogger,
	serviceRepo *domain.IServiceRepo) (m *SwaggerModel) {
	m = &SwaggerModel{}

	if coll == nil || log == nil || serviceRepo == nil {
		return nil
	}
	m.coll = &coll
	m.log = log
	m.serviceRepo = serviceRepo
	return
}

// Implements ISwaggerRepo method
func (o SwaggerModel) CreateSwagger(d *domain.SwaggerCreateData) (string, error) {
	now := primitive.NewDateTimeFromTime(time.Now())

	db := domain.SwaggerModel{
		Version:    d.Version,
		ServiceRef: d.ServiceRef,
		OwnerRef:   d.OwnerRef,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	if d.SourceType == "doc" {
		db.SwaggerSpec = d.Source
	} else if d.SourceType == "url" {
		return "", errors.New("ERR_SWAGGER_TODO")
	} else {
		return "", errors.New("ERR_SWAGGER_TODO")
	}

	// TODO: add some checks for the incoming swagger
	var openApiV3 domain.OpenApiV3
	byteDoc, _ := json.Marshal(db.SwaggerSpec)
	_ = json.Unmarshal(byteDoc, &openApiV3)
	db.Endpoints, _ = openApiV3.GetEndPts()

	res, err := (*o.coll).InsertOne(context.TODO(), db)
	if err != nil {
		(*o.log).Errorln("ERR_SWAGGER_TODO", err.Error())
		return "", err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		(*o.log).Errorln("ERR_SWAGGER_TODO", err.Error())
		return "<internal_error>", nil
	}

	err = (*o.serviceRepo).UpdateService(d.ServiceRef, &domain.ServiceUpdateData{
		SwaggerRef: id,
	})
	if err != nil {
		// ignore errors while updating, just log it
		(*o.log).Errorln("ERR_SWAGGER_TODO", err.Error())
	}

	return id.Hex(), nil
}

// Implements ISwaggerRepo method
func (o SwaggerModel) GetSwaggers(d *domain.SwaggerQueryData) ([]domain.SwaggerModel, error) {
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
		(*o.log).Errorln("ERR_SWAGGER_TODO", err.Error())
		return nil, err
	}

	res := make([]domain.SwaggerModel, 0)
	if err = cursor.All(context.TODO(), &res); err != nil {
		(*o.log).Errorln("ERR_SWAGGER_TODO", err.Error())
		return nil, err
	}
	return res, nil
}

// Implements ISwaggerRepo method
func (o SwaggerModel) GetOneSwagger(d *domain.SwaggerQueryData) (*domain.SwaggerModel, error) {
	if d.Id == primitive.NilObjectID {
		return nil, errors.New("ERR_SWAGGER_TODO")
	}

	filter := map[string]any{
		"_id": d.Id,
	}

	var res domain.SwaggerModel
	if err := (*o.coll).FindOne(context.TODO(), filter).Decode(&res); err != nil {
		return nil, err
	} else {
		return &res, nil
	}
}

// Implements ISwaggerRepo method
func (o SwaggerModel) UpdateSwagger(d *domain.SwaggerUpdateData) error {
	return nil
}

func (o SwaggerModel) getEndPoints(s any) []domain.SwaggerEndPointMetadata {
	m := make([]domain.SwaggerEndPointMetadata, 0)
	spec, ok := s.(*domain.OpenApiV3)
	if !ok {
		fmt.Println("not ok")
		return m
	}

	res, err := spec.GetEndPts()
	if err != nil {
		fmt.Println(err.Error())
		return m
	}

	return res
}
