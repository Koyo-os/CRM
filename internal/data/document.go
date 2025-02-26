package data

import (
	"context"
	"time"

	"github.com/koyo-os/crm/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *DocRepository) GetDocument(id uint64) (*models.Document, error) {
	var doc models.Document

	res := r.coll.FindOne(r.ctx, bson.M{
		"id" : id,
	})

	err := res.Decode(&doc)
	return &doc, err
}

