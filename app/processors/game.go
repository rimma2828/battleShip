package processors

import (
	"go.uber.org/zap"

	"battleship/app/repository"
	"battleship/app/system/db"
)

type GameProcessor interface {
	Reset() (bool, error)
}

type gameProcessor struct {
	client         db.Storage
	repoFight      repository.FightRepository
	repoCoordinate repository.CoordinatesRepository
	log            *zap.Logger
}

func NewGameProcessor(client db.Storage, repoFight repository.FightRepository, repoCoordinate repository.CoordinatesRepository, log *zap.Logger) GameProcessor {
	return &gameProcessor{
		client:         client,
		repoFight:      repoFight,
		repoCoordinate: repoCoordinate,
		log:            log,
	}
}

func (c *gameProcessor) Reset() (bool, error) {
	storage := c.client
	status, err := c.repoCoordinate.Truncate(storage)
	if err != nil {
		return false, err
	}
	status, err = c.repoFight.Truncate(storage)
	if err != nil {
		return false, err
	}

	return status, nil
}
