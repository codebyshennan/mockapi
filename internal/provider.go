package internal

import (
	"context"
	"time"

	"github.com/codebyshennan/mockapi/domain"
	"github.com/codebyshennan/mockapi/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// A struct containing all the repositories,
// to be injected.
type Provider struct {
	UserRepo            domain.IUserRepo
	MockServerRepo      domain.IMockServerRepo
	MockEndPtRepo       domain.IMockEndPtRepo
	MockServerStoreRepo domain.IMockServerStoreRepo
	ServiceRepo         domain.IServiceRepo
	SwaggerRepo         domain.ISwaggerRepo
	AuthRepo            domain.IAuthRepo
	Config              *domain.Config
	Logger              domain.ILogger

	Db *mongo.Database
}

func NewProvider() (p *Provider) {
	p = &Provider{}

	// Read in configs/env variables
	if env, err := ReadEnv(); err != nil {
		return nil
	} else {
		p.Config = env
	}

	// Initialise db connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(p.Config.MongoUrl))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// Load all dependencies to be injected
	p.Logger = getLogger()

	p.Db = client.Database(p.Config.DbName)

	p.UserRepo = models.NewUserModel(
		p.Db.Collection("users"), p.Logger)
	p.AuthRepo = models.NewAuthModel(p.UserRepo, p.Logger, p.Config)

	p.ServiceRepo = models.NewServiceModel(
		p.Db.Collection("services"), p.Logger)
	p.SwaggerRepo = models.NewSwaggerModel(
		p.Db.Collection("services_swaggers"), &p.Logger, &p.ServiceRepo)

	p.MockEndPtRepo = models.NewMockEndPtModel(
		p.Db.Collection("mockServers_mockEndPts"))
	p.MockServerRepo = models.NewMockServerModel(
		p.Db.Collection("mockServers"), p.MockEndPtRepo)

	p.MockServerStoreRepo = models.NewMockServerStoreRepo(
		p.Db.Collection("mockServers_stores"), p.Logger)

	return
}
