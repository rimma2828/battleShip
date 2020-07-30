package repository

import (
	"context"

	gen "battleship/app/generated/models"
	"battleship/app/models"
	"battleship/app/system/db"
)

type CoordinatesRepositoryMock struct {
	SaveResponseFn             func(ctx context.Context) (bool, error)
	TruncateResponseFn         func() (bool, error)
	GetByCoordinatesResponseFn func(ctx context.Context) (*models.Coordinates, error)
}

func (c *CoordinatesRepositoryMock) Save(client db.Storage, coords *gen.CoordinatesAddParams) (bool, error) {
	return c.SaveResponseFn(context.WithValue(context.Background(), "coords", coords))
}

func (c *CoordinatesRepositoryMock) Truncate(client db.Storage) (bool, error) {
	return c.TruncateResponseFn()
}

func (c *CoordinatesRepositoryMock) GetByCoordinates(client db.Storage, x_coord int64, y_coord int64, user_id int64) (*models.Coordinates, error) {
	return c.GetByCoordinatesResponseFn(context.WithValue(context.WithValue(context.WithValue(context.Background(), "user_id", user_id), "y_coord", y_coord), "x_coord", x_coord))
}
