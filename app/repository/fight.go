package repository

import (
	"battleship/app/errors"
	gen "battleship/app/generated/models"
	"battleship/app/system/db"
)

type FightRepository interface {
	Save(client db.Storage, params *gen.FightActParams, shipId int64) (bool, error)
	Truncate(client db.Storage) (bool, error)
}

type fightRepository struct {
}

func NewFightRepository() FightRepository {
	return &fightRepository{}
}

func (c *fightRepository) Save(client db.Storage, params *gen.FightActParams, shipId int64) (bool, error) {
	q := `insert into fight_act (user_id, x_coord, y_coord, ship_id) 
			VALUES ($1, $2, $3, $4);`
	err := client.Exec(
		q,
		params.UserID,
		params.XCoord,
		params.YCoord,
		shipId,
	)
	if err != nil {
		return false, errors.NewDatabaseError(errors.DatabaseBasicError, err)
	}

	return true, nil
}

func (c *fightRepository) Truncate(client db.Storage) (bool, error) {
	q := `truncate table fight_act;`
	err := client.Exec(q)
	if err != nil {
		return false, errors.NewDatabaseError(errors.DatabaseBasicError, err)
	}

	return true, nil
}
