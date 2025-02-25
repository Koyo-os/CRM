package data

import (
	"context"
	"math/rand"
	"time"

	"github.com/koyo-os/crm/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *UserRepository) AddUser(user *models.User) (uint64, error) {
	user.ID = rand.Uint64()
	_, err := r.coll.InsertOne(r.ctx, user)
	return user.ID, err
}

func (r *UserRepository) DeleteUser(id uint64) error {
	res := r.coll.FindOneAndDelete(r.ctx, bson.M{
		"id" : id,
	})

	return res.Err()
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	var user models.User

	cursor, err := r.coll.Find(r.ctx, bson.M{})
	if err != nil{
		return nil, err
	}

	for cursor.Next(r.ctx) {
		cursor.Decode(&user)
		users = append(users, user)
	}

	if cursor.Err() != nil{
		return nil, err
	}

	return users, nil
}
func (r *UserRepository) CheckSuperUser(ID uint64, key string) (bool, error) {
	var result bson.M
    err := r.coll.FindOne(r.ctx, bson.M{
		"id" : ID,
		"key" : key,
	}).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return false, err
        } else {
			return false, err
        }
    } else {
		return true, nil
	}
}

func (r *UserRepository) GetUser(id uint64) (*models.User, error) {
	filter := bson.M{
		"id" : id,
	}

	var user models.User

	res := r.coll.FindOne(r.ctx, filter)
	err := res.Decode(&user)
	return &user, err
}
