package game

import (
	"context"

	respBuilders "battleship/app/builders/responses"
	"battleship/app/errors"
)

func (c *Game) Reset(ctx context.Context, params interface{}) (interface{}, error) {
	status, err := c.processor.Reset()
	if err != nil {
		switch err.(type) {
		case errors.DatabaseError:
			c.log.Error(err.Error())
		default:
			c.log.Warn(err.Error())
		}
		return nil, errors.WrapToJsonRpcError(err)
	}
	return respBuilders.BuildCoordinatesActResponse(status, "Новая игра"), nil
}
