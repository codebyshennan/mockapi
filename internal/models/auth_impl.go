package models

import (
	"context"
	"errors"
	"time"

	"github.com/codebyshennan/mockapi/domain"
	"github.com/codebyshennan/mockapi/internal/common/h"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/idtoken"
)

type AuthModel struct {
	userRepo domain.IUserRepo
	logger   domain.ILogger
	config   *domain.Config
}

func NewAuthModel(userRepo domain.IUserRepo, logger domain.ILogger, config *domain.Config) (m AuthModel) {
	m = AuthModel{
		logger:   logger,
		userRepo: userRepo,
		config:   config,
	}
	return
}

func (o AuthModel) GoogleLogin(d *domain.GoogleLoginPostData) (*domain.AuthRes, error) {
	payload, err := idtoken.Validate(context.Background(), d.Credential, o.config.GoogleClientId)
	if err != nil {
		o.logger.Errorln("ERR_AUTH_GOOGLE-TODO", err.Error())
		return nil, err
	}

	var createData = domain.UserCreateData{
		FirstName: payload.Claims["given_name"].(string),
		LastName:  payload.Claims["family_name"].(string),
		Email:     payload.Claims["email"].(string),
		IsRoot:    false,
	}

	user, err := o.userRepo.GetOrInsert(&createData)
	if err != nil {
		o.logger.Errorln(h.MongoError, err.Error())
		return nil, err
	}

	token, err := o.generateJWT(user.Id.Hex())
	if err != nil {
		o.logger.Errorln("ERR_AUTH_TODO", err.Error())
		return nil, err
	}

	return &domain.AuthRes{
		Token:     token,
		UserModel: *user,
	}, nil
}

// GenerateJWT generates a JSON Web Token signed using the env variable JWT_SIGNING_KEY
func (o AuthModel) generateJWT(id string) (signedToken string, err error) {
	mySigningKey := []byte(o.config.JwtKey)

	// Create the claims
	claims := domain.JWTClaims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(), // 24H expiry
		},
	}

	// Create the token and sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString(mySigningKey)
	if err != nil {
		o.logger.Errorln("ERR_AUTH_TODO", err)
		return
	}

	return
}

func (o AuthModel) ValidateAndDecodeJWT(token string) (user *domain.UserModel, err error) {
	c := domain.JWTClaims{}

	// Contains in built check to see if token has expired
	_, err = jwt.ParseWithClaims(token, &c, func(token *jwt.Token) (interface{}, error) {
		return []byte(o.config.JwtKey), nil
	})

	if err != nil {
		o.logger.Errorln("ERR_AUTH_PARSE-JWT", err)
		return nil, errors.New("ERR_AUTH_TODO")
	}

	oid, err := primitive.ObjectIDFromHex(c.UserID)
	if err != nil {
		o.logger.Errorln("ERR_AUTH_OID", err)
		return nil, errors.New("ERR_AUTH_TODO")
	}

	user, err = o.userRepo.GetOneUser(&domain.UserQueryData{
		Id: oid,
	})
	if err != nil {
		o.logger.Errorln("ERR_AUTH_GET-USER", err)
		return nil, errors.New("ERR_AUTH_TODO")
	}

	return user, nil
}
