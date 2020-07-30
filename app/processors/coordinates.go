package processors

import (
	gen "battleship/app/generated/models"
	"battleship/app/repository"
	"battleship/app/system/db"

	"go.uber.org/zap"
)

type CoordinatesProcessor interface {
	Add(params *gen.CoordinatesAddParams) (bool, error)
}

type coordinatesProcessor struct {
	client          db.Storage
	repoCoordinates repository.CoordinatesRepository
	repoUser        repository.UserRepository
	repoShip        repository.ShipRepository
	log             *zap.Logger
}

func NewCoordinatesProcessor(client db.Storage, repoCoordinates repository.CoordinatesRepository, repoUser repository.UserRepository, repoShip repository.ShipRepository, log *zap.Logger) CoordinatesProcessor {
	return &coordinatesProcessor{
		client:          client,
		repoCoordinates: repoCoordinates,
		repoUser:        repoUser,
		repoShip:        repoShip,
		log:             log,
	}
}

func (c *coordinatesProcessor) Add(params *gen.CoordinatesAddParams) (bool, error) {
	storage := c.client
	_, err := c.repoUser.Get(storage, params.UserID)
	if err != nil {
		return false, err
	}
	_, err = c.repoShip.Get(storage, params.ShipID)
	if err != nil {
		return false, err
	}
	return c.repoCoordinates.Save(storage, params)
}
