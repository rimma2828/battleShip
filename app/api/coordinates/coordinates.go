package coordinates

import (
	"go.uber.org/zap"

	"battleship/app/processors"
)

type Coordinates struct {
	processor processors.CoordinatesProcessor
	log       *zap.Logger
}

func NewCoordinates(
	processor processors.CoordinatesProcessor,
	log *zap.Logger,
) *Coordinates {
	return &Coordinates{
		processor: processor,
		log:       log,
	}
}
