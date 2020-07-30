package processors

import (
	"fmt"
	"go.uber.org/zap"

	"battleship/app/models"
	"battleship/app/repository"
	"battleship/app/system/db"
)

type UserProcessor interface {
	Get(id int64) (*models.User, error)
	Create(user *models.User) (*models.User, error)
}

type userProcessor struct {
	client   db.Storage
	repoUser repository.UserRepository
	log      *zap.Logger
}

func NewUserProcessor(client db.Storage, repoUser repository.UserRepository, log *zap.Logger) UserProcessor {
	return &userProcessor{
		client:   client,
		repoUser: repoUser,
		log:      log,
	}
}

func (c *userProcessor) Get(id int64) (*models.User, error) {
	user, err := c.repoUser.Get(c.client, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *userProcessor) Create(user *models.User) (*models.User, error) {
	var id int64
	storage := c.client

	respUser, err := c.repoUser.GetByName(storage, user.Name)
	if respUser != nil {
		return nil, fmt.Errorf("пользователь с таким именем уже существует")
	}
	id, err = c.repoUser.Save(storage, user)
	if err != nil {
		return nil, err
	}
	respUser, err = c.repoUser.Get(storage, id)
	if err != nil {
		return nil, err
	}

	return respUser, nil
}
