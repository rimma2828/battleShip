package processors

import (
	"go.uber.org/zap"

	"battleship/app/models"
	"battleship/app/repository"
	"battleship/app/system/db"
)

type ShipProcessor interface {
	GetList() ([]models.Ship, error)
}

type shipProcessor struct {
	client   db.Storage
	repoShip repository.ShipRepository
	log      *zap.Logger
}

func NewShipProcessor(client db.Storage, repoShip repository.ShipRepository, log *zap.Logger) ShipProcessor {
	return &shipProcessor{
		client:   client,
		repoShip: repoShip,
		log:      log,
	}
}

func (c *shipProcessor) GetList() ([]models.Ship, error) {
	shipList, err := c.repoShip.GetList(c.client)
	if err != nil {
		return nil, err
	}

	return shipList, nil
}
