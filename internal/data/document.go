package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type DocRepository struct{
	coll *mongo.Collection
	ctx context.Context
}

func NewDocRepo(client *mongo.Client) *DocRepository{
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	return &DocRepository{
		coll: client.Database("crm").Collection("user"),
		ctx: ctx,
	}
}