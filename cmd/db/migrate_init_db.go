package main

import (
	"context"
	"fmt"

	"github.com/codebyshennan/mockapi/domain"
	"github.com/codebyshennan/mockapi/internal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type serviceDb struct {
	Id          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"serviceName"`
	SwaggerRefs []primitive.ObjectID `bson:"swaggers"`
}

type swaggerDb struct {
	Id        primitive.ObjectID `bson:"_id"`
	OpenApiV3 map[string]any     `bson:"openApiV3"`
	Version   string             `bson:"version"`
}

// Migrate previous staging DB into this one
func main() {
	provider := internal.NewProvider()
	if provider == nil {
		fmt.Println("init error")
	}

	// fetch all services in old db
	svcColl := provider.Db.Collection("svcs_1")
	cursor, _ := svcColl.Find(context.Background(), map[string]any{})
	svcs := make([]serviceDb, 0)
	if err := cursor.All(context.TODO(), &svcs); err != nil {
		return
	}

	// fetch all swaggers in old db
	swaggerColl := provider.Db.Collection("swaggers_1")
	cursor, _ = swaggerColl.Find(context.Background(), map[string]any{})
	swaggers := make([]swaggerDb, 0)
	if err := cursor.All(context.TODO(), &swaggers); err != nil {
		return
	}

	for _, svc := range svcs {
		if len(svc.SwaggerRefs) == 0 {
			continue
		}

		if len(svc.SwaggerRefs) > 1 {
			fmt.Println(svc.Id, " has more than 1 swaggerRef")
		}

		id := svc.SwaggerRefs[0]
		persId, err := primitive.ObjectIDFromHex("63113b9c2644cc400660ac83")
		if err != nil {
			return
		}

		for _, swagger := range swaggers {
			if swagger.Id == id {
				fmt.Println("Matched ", swagger.Id)
				nid, err := provider.ServiceRepo.CreateService(&domain.ServiceCreateData{
					Name:        svc.Name,
					Description: "",
				})
				if err != nil {
					fmt.Println("Error while creating (svc) ", nid)
				}
				_, err = provider.SwaggerRepo.CreateSwagger(&domain.SwaggerCreateData{
					OwnerRef:   persId,
					Version:    swagger.Version,
					SourceType: "doc",
					Source:     swagger.OpenApiV3,
					ServiceRef: nid,
				})
				if err != nil {
					fmt.Println("Error while creating (swagger) ", nid)
				}
			}
		}
	}

}
