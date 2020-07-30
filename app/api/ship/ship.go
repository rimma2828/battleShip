package ship

import (
	"go.uber.org/zap"

	"battleship/app/processors"
)

type Ship struct {
	processor processors.ShipProcessor
	log       *zap.Logger
}

func NewShip(
	processor processors.ShipProcessor,
	log *zap.Logger,
) *Ship {
	return &Ship{
		processor: processor,
		log:       log,
	}
}
