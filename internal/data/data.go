package data

import (
	"context"
	"time"

	"github.com/koyo-os/crm/internal/config"
	"github.com/koyo-os/crm/internal/data/models"
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

func (r *Repository) CheckDocOnUserPermision(userID, docID uint64, typeRope rune, number uint8) (bool, error) {
	user, err := r.User.GetUser(userID)
	if err != nil{
		return false, err
	}

	doc, err := r.Docs.GetDocument(docID)
	if err != nil{
		return false, err
	}

	ok := false
	for _, p := range doc.Roles{
		for _, d := range user.Role {
			if p == d.Name && d.TypeRole[number] == typeRope{
				ok = true
			}
		}
	}

	return ok, nil
}
 
func (r *Repository) GetDocsByUserPermitions(userID uint64) ([]models.Document, error) {
	var res []models.Document

	docs, err := r.Docs.GetAll()
	if err != nil{
		return nil, err
	}

	user, err := r.User.GetUser(userID)
	if err != nil{
		return nil,err
	}

	for _, d := range docs {
		can := false
		for _, r := range user.Role {
			for _, j := range d.Roles {
				if r.Name == j {
					can = true
				}
			}
		}
		if can {
			res = append(res, d)
		}
	}

	return res, nil
}