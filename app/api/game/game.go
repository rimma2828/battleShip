package game

import (
	"go.uber.org/zap"

	"battleship/app/processors"
)

type Game struct {
	processor processors.GameProcessor
	log       *zap.Logger
}

func NewGame(
	processor processors.GameProcessor,
	log *zap.Logger,
) *Game {
	return &Game{
		processor: processor,
		log:       log,
	}
}
