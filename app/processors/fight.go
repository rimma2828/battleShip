package processors

import (
	"go.uber.org/zap"

	gen "battleship/app/generated/models"
	"battleship/app/repository"
	"battleship/app/system/db"
)

type FightProcessor interface {
	Act(params *gen.FightActParams) (bool, string, error)
}

const away = "мимо"
const shot = "ранен"

type fightProcessor struct {
	client    db.Storage
	repoFight repository.FightRepository
	repoUser  repository.UserRepository
	repoCoord repository.CoordinatesRepository
	log       *zap.Logger
}

func NewFightProcessor(client db.Storage, repoFight repository.FightRepository, repoUser repository.UserRepository, repoCoord repository.CoordinatesRepository, log *zap.Logger) FightProcessor {
	return &fightProcessor{
		client:    client,
		repoFight: repoFight,
		repoUser:  repoUser,
		repoCoord: repoCoord,
		log:       log,
	}
}

func (c *fightProcessor) Act(params *gen.FightActParams) (bool, string, error) {
	storage := c.client
	_, err := c.repoUser.Get(storage, params.UserID)
	if err != nil {
		return false, "", err
	}
	coords, err := c.repoCoord.GetByCoordinates(c.client, params.XCoord, params.YCoord, params.UserID)
	if err != nil {
		return false, "", err
	}
	var shipId int64
	shipId = 0
	if coords != nil {
		shipId = coords.ShipId
	}
	status, err := c.repoFight.Save(storage, params, shipId)
	if err != nil {
		return false, "", err
	}
	if coords != nil {
		return status, shot, nil
	} else {
		return status, away, nil
	}
}
