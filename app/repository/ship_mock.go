package repository

import (
	"context"

	"battleship/app/models"
	"battleship/app/system/db"
)

type ShipRepositoryMock struct {
	GetListResponseFn func() ([]models.Ship, error)
	GetResponseFn     func(ctx context.Context) (*models.Ship, error)
}

func (c *ShipRepositoryMock) Get(client db.Storage, id int64) (*models.Ship, error) {
	return c.GetResponseFn(context.WithValue(context.Background(), "id", id))
}

func (c *ShipRepositoryMock) GetList(client db.Storage) ([]models.Ship, error) {
	return c.GetListResponseFn()
}
