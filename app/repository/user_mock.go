package repository

import (
	"context"

	"battleship/app/models"
	"battleship/app/system/db"
)

type UserRepositoryMock struct {
	GetResponseFn       func(ctx context.Context) (*models.User, error)
	GetByNameResponseFn func(ctx context.Context) (*models.User, error)
	SaveResponseFn      func(ctx context.Context) (int64, error)
}

func (c *UserRepositoryMock) Get(client db.Storage, id int64) (*models.User, error) {
	return c.GetResponseFn(context.WithValue(context.Background(), "id", id))
}

func (c *UserRepositoryMock) GetByName(client db.Storage, name string) (*models.User, error) {
	return c.GetByNameResponseFn(context.WithValue(context.Background(), "name", name))
}

func (c *UserRepositoryMock) Save(client db.Storage, user *models.User) (int64, error) {
	return c.SaveResponseFn(context.WithValue(context.Background(), "user", user))
}
