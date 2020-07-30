package repository

import (
	"fmt"
	"battleship/app/errors"
	appModels "battleship/app/models"
	"battleship/app/system/db"
)

type ShipRepository interface {
	GetList(client db.Storage) ([]appModels.Ship, error)
	Get(client db.Storage, id int64) (*appModels.Ship, error)
}

type shipRepository struct {
}

func NewShipRepository() ShipRepository {
	return &shipRepository{}
}

func (c *shipRepository) GetList(client db.Storage) ([]appModels.Ship, error) {
	ships := []appModels.Ship{}

	err := client.Query(&ships,
		`select 
		 id, 
		 length,
		 count
		from ship;`)

	if err != nil {
		return nil, errors.NewDatabaseError(errors.DatabaseBasicError, err)
	}

	return ships, nil
}

func (c *shipRepository) Get(client db.Storage, id int64) (*appModels.Ship, error) {
	var ship []appModels.Ship

	err := client.Query(&ship,
		`select 
		 id, 
		 length,
		 count
		from ship where id = $1
		limit 1;`,
		id)

	if err != nil {
		return nil, errors.NewDatabaseError(errors.DatabaseBasicError, err)
	}
	if len(ship) == 0 {
		return nil, fmt.Errorf("корабль не найден")
	}

	return &ship[0], nil
}
