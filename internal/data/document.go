package data

import (
	"context"
	"math/rand/v2"
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

func (r *DocRepository) AddDocument(doc *models.Document) (uint64, error) {
	doc.ID = rand.Uint64()
	_, err := r.coll.InsertOne(r.ctx, doc)
	return doc.ID, err
}

func (r *DocRepository) GetAll() ([]models.Document, error) {
	var docs []models.Document
	var doc models.Document

	cursor, err := r.coll.Find(r.ctx, bson.M{})
	if err != nil{
		return nil, err
	}

	for cursor.Next(r.ctx) {
		cursor.Decode(&doc)
		docs = append(docs, doc)
	}

	return docs, cursor.Err()
}

func (r *DocRepository) Delete(id uint64) error {
	filter := bson.M{
		"id" : id,
	}

	res :=  r.coll.FindOneAndDelete(r.ctx, filter)
	return res.Err()
}

func (r *DocRepository) Update(id uint64, newDocs *models.Document) error {
	filter := bson.M{
		"id" : id,
	}

	update := bson.M{
		"content" : newDocs.Content,
		"about" : newDocs.About,
		"roles" : newDocs.Roles,
	}

	res := r.coll.FindOneAndUpdate(r.ctx, filter, update)
	return res.Err()
}