package models

import (
	"context"
	"errors"
	"time"

	"github.com/codebyshennan/mockapi/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// An actual instance of user model tied to the db
type UserModel struct {
	coll domain.ICollection
	log  domain.ILogger
}

// Factory method to create user model
func NewUserModel(coll domain.ICollection, log domain.ILogger) (m *UserModel) {
	m = &UserModel{
		coll: coll,
		log:  log,
	}

	if coll == nil || log == nil {
		return nil
	}
	return
}

// Implements IUserRepo method
func (o UserModel) CreateUser(d *domain.UserCreateData) (string, error) {
	now := primitive.NewDateTimeFromTime(time.Now())

	db := domain.UserModel{
		FirstName: d.FirstName,
		LastName:  d.LastName,
		Email:     d.Email,
		IsRoot:    false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	res, err := o.coll.InsertOne(context.TODO(), db)
	if err != nil {
		o.log.Errorln("ERR_USER_TODO", err.Error())
		return "", err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		o.log.Errorln("ERR_USER_TODO", err.Error())
		return "<internal_error>", nil
	}

	return id.Hex(), nil
}

// Implements IUserRepo method
func (o UserModel) GetUsers(d *domain.UserQueryData) ([]domain.UserModel, error) {
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

	cursor, err := o.coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		o.log.Errorln("ERR_USER_TODO", err.Error())
		return nil, err
	}

	res := make([]domain.UserModel, 0)
	if err = cursor.All(context.TODO(), &res); err != nil {
		o.log.Errorln("ERR_USER_TODO", err.Error())
		return nil, err
	}
	return res, nil
}

// Implements IUserRepo method
func (o UserModel) GetOneUser(d *domain.UserQueryData) (*domain.UserModel, error) {
	if d.Id == primitive.NilObjectID && d.Email == "" {
		return nil, errors.New(domain.UserNotFound)
	}

	filter := map[string]any{}
	if d.Id != primitive.NilObjectID {
		filter["_id"] = d.Id
	}
	if d.Email != "" {
		filter["email"] = d.Email
	}

	var res domain.UserModel
	if err := o.coll.FindOne(context.TODO(), filter).Decode(&res); err != nil {
		o.log.Errorln(domain.UserNotFound, err.Error())
		return nil, errors.New(domain.UserNotFound)
	} else {
		return &res, nil
	}
}

// Implements IUserRepo method
func (o UserModel) UpdateUser(d *domain.UserUpdateData) error {
	return nil
}

// Implements IUserRepo method
func (o UserModel) GetOrInsert(d *domain.UserCreateData) (*domain.UserModel, error) {
	user, err := o.GetOneUser(&domain.UserQueryData{
		Email: d.Email,
	})

	// user is found
	if err == nil {
		return user, nil
	}

	// user is not found
	if err != nil && err.Error() == domain.UserNotFound {
		id, err := o.CreateUser(d)
		if err != nil {
			return nil, err
		}

		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}

		user, err := o.GetOneUser(&domain.UserQueryData{
			Id: oid,
		})
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	// other errors
	return nil, err
}
