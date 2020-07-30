package repository

import (
	"battleship/app/errors"
	gen "battleship/app/generated/models"
	appModels "battleship/app/models"
	"battleship/app/system/db"
)

type CoordinatesRepository interface {
	Save(client db.Storage, coords *gen.CoordinatesAddParams) (bool, error)
	Truncate(client db.Storage) (bool, error)
	GetByCoordinates(client db.Storage, x_coord int64, y_coord int64, user_id int64) (*appModels.Coordinates, error)
}

type coordinatesRepository struct {
}

func NewCoordinatesRepository() CoordinatesRepository {
	return &coordinatesRepository{}
}

func (c *coordinatesRepository) Save(client db.Storage, coords *gen.CoordinatesAddParams) (bool, error) {
	q := `insert into coords_user_ship (user_id, ship_id, x_coord, y_coord) 
			VALUES ($1, $2, $3, $4);`

	err := client.Exec(
		q,
		coords.UserID,
		coords.ShipID,
		coords.XCoord,
		coords.YCoord,
	)
	if err != nil {
		return false, errors.NewDatabaseError(errors.DatabaseBasicError, err)
	}

	return true, nil
}

func (c *coordinatesRepository) Truncate(client db.Storage) (bool, error) {
	q := `truncate table coords_user_ship;`
	err := client.Exec(q)
	if err != nil {
		return false, errors.NewDatabaseError(errors.DatabaseBasicError, err)
	}

	return true, nil
}

func (c *coordinatesRepository) GetByCoordinates(client db.Storage, x_coord int64, y_coord int64, user_id int64) (*appModels.Coordinates, error) {
	var coords []appModels.Coordinates

	err := client.Query(&coords,
		`select 
		 user_id,
		 ship_id, 
		 x_coord,
		 y_coord
		from coords_user_ship
		where x_coord = $1 and y_coord = $2 and user_id != $3 limit 1;`,
		x_coord,
		y_coord,
		user_id)

	if err != nil {
		return nil, errors.NewDatabaseError(errors.DatabaseBasicError, err)
	}
	if len(coords) == 0 {
		return nil, nil
	}

	return &coords[0], nil
}
