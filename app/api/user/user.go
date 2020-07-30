package user

import (
	"go.uber.org/zap"

	"battleship/app/processors"
)

type User struct {
	processor processors.UserProcessor
	log       *zap.Logger
}

func NewUser(
	processor processors.UserProcessor,
	log *zap.Logger,
) *User {
	return &User{
		processor: processor,
		log:       log,
	}
}
