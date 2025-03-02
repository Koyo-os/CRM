package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/koyo-os/crm/internal/config"
	"github.com/koyo-os/crm/internal/data"
	"github.com/koyo-os/crm/internal/data/models"
	"github.com/koyo-os/crm/pkg/loger"
	"golang.org/x/crypto/bcrypt"
)

type Service struct{
	logger loger.Logger
	Repo *data.Repository
}

func New(cfg *config.Config) (*Service, error) {
	repo, err := data.New(cfg)

	return &Service{
		Repo: repo,
		logger: loger.New(),
	}, err
}

func (s *Service) GetDocument(Userid, docID uint64,key string) (*models.Document, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil{
		return nil,err
	}

	ok, err := s.Repo.User.CheckUser(Userid, string(hash))
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
	if err != nil{
		return err
	}
	
	if permsOk && ok {
		return s.Repo.Docs.Delete(docid)
	}

	return errors.New("you dont have permition to do it")
}

func generateJwt(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"id" : user.ID,
		"key" : user.Key,
		"exp" : time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SEKRET_KEY")))
}

func (s *Service) CreateUser(user *models.User) (string,uint64, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Key), bcrypt.DefaultCost)
	if err != nil{
		return "", 0, err
	}

	user.Key = string(hash)
	
	id, err := s.Repo.User.AddUser(user)
	if err != nil{
		return "", 0, err
	}

	token, err := generateJwt(user)
	if err != nil{
		return "", 0, err
	}

	return token, id, nil
}