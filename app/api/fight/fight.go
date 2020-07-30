package fight

import (
	"go.uber.org/zap"

	"battleship/app/processors"
)

type Fight struct {
	processor processors.FightProcessor
	log       *zap.Logger
}

func NewFight(
	processor processors.FightProcessor,
	log *zap.Logger,
) *Fight {
	return &Fight{
		processor: processor,
		log:       log,
	}
}
