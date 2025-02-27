package service

import (
	"errors"
	"time"

	"github.com/koyo-os/crm/internal/data"
	"github.com/koyo-os/crm/internal/data/models"
)

type Service struct{
	Repo *data.Repository
}

func New(repo *data.Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) GetDocument(Userid, docID uint64,key string) (*models.Document, error) {
	ok, err := s.Repo.User.CheckUser(Userid, key)
	if err != nil{
		return nil, err
	}

	okPerms,err := s.Repo.CheckDocOnUserPermision(Userid, docID)
	if err != nil{
		return nil, err
	}

	if ok {
		return s.Repo.Docs.GetDocument(docID)
	}

	if ok && okPerms {
		return s.Repo.Docs.GetDocument(docID)
	}

	return nil, errors.New("you dont have permitions for this doc")
}

func (s *Service) CheckAllUserRoleTimes() error {
	now := time.Now().Format(models.TIME_LAYOUT)

	users, err := s.Repo.User.GetUsers()
	if err != nil{
		return err
	}

	for _, u := range users {
		for _, r := range u.Role {
			
		}
	}
}