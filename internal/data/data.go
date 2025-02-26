package data

import (
	"context"
	"time"

	"github.com/koyo-os/crm/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct{
	User *UserRepository
	Docs *DocRepository
}

func New(cfg config.Config) (*Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	url := options.Client().ApplyURI(cfg.MongoURL)
	client, err := mongo.Connect(ctx, url)
	if err != nil{
		return nil, err
	}

	return &Repository{
		User: NewUserRepo(client),
		Docs: NewDocRepo(client),
	}, nil
}

func (r *Repository) CheckDocOnUserPermision(userID, docID uint64) (bool, error) {
	user, err := r.User.GetUser(userID)
	if err != nil{
		return false, err
	}

	
}
 
