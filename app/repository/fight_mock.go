package repository

import (
	"context"
	gen "battleship/app/generated/models"
	"battleship/app/system/db"
)

type FightRepositoryMock struct {
	SaveResponseFn     func(ctx context.Context) (bool, error)
	TruncateResponseFn func() (bool, error)
}

func (c *FightRepositoryMock) Save(client db.Storage, fight *gen.FightActParams, shipId int64) (bool, error) {
	return c.SaveResponseFn(context.WithValue(context.Background(), "fight", fight))
}

func (c *FightRepositoryMock) Truncate(client db.Storage) (bool, error) {
	return c.TruncateResponseFn()
}
