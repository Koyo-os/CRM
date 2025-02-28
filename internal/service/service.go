package service

import (
	"errors"
	"time"

	"github.com/koyo-os/crm/internal/data"
	"github.com/koyo-os/crm/internal/data/models"
	"github.com/koyo-os/crm/pkg/loger"
)

type Service struct{
	logger loger.Logger
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

	okPerms,err := s.Repo.CheckDocOnUserPermision(Userid, docID, 'g', 0)
	if err != nil{
		return nil, err
	}

	if ok && okPerms {
		return s.Repo.Docs.GetDocument(docID)
	}

	return nil, errors.New("you dont have permitions for this doc")
}

func (s *Service) AddDocument(UserID uint64,key string, doc *models.Document) (uint64, error) {
	ok, err := s.Repo.User.CheckUser(UserID, key)
	if err != nil{
		return 0, err
	}

	user, err := s.Repo.User.GetUser(UserID)
	if err != nil{
		return 0, err
	}

	can := false 
	for _, v := range user.Role {
		if v.CanAddDoc {
			can = true
		}
	}

	if can && ok{
		return s.Repo.Docs.AddDocument(doc)
	} else {
		return 0, errors.New("permition denied")
	}
}

func (s *Service) CheckAllUserRoleTimes(ch chan error) {
	now := time.Now().Format(models.TIME_LAYOUT)

	users, err := s.Repo.User.GetUsers()
	if err != nil{
		ch <- err
		return
	}

	s.logger.Info().Msg("starting role timeout check!")

	for _, u := range users {
		for _, r := range u.Role {
			if r.TimeToEnd.Format(models.TIME_LAYOUT) == now {
				if err := s.Repo.User.DeleteUserRole(u.ID, r.Name);err != nil{
					s.logger.Error().Err(err)
				}
			}
		}
	}

}

func (s *Service) DeleteDocument(docid, userid uint64, key string) error {
	ok, err := s.Repo.User.CheckUser(userid, key)
	if err != nil{
		return err
	}

	permsOk, err := s.Repo.CheckDocOnUserPermision(userid, docid, 'd', 1)
	if permsOk && ok {
		return s.Repo.Docs.
	}
}