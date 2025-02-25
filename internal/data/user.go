package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct{
	coll *mongo.Collection
	ctx context.Context
}

func NewUserRepo(client *mongo.Client) *UserRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	return &UserRepository{
		coll: client.Database("crm").Collection("user"),
		ctx: ctx,
	}
}