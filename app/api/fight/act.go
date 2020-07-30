package fight

import (
	"context"

	respBuilders "battleship/app/builders/responses"
	"battleship/app/errors"
	gen "battleship/app/generated/models"
)

func (c *Fight) Act(ctx context.Context, params gen.FightActParams) (interface{}, error) {
	status, shipStatus, err := c.processor.Act(&params)
	if err != nil {
		switch err.(type) {
		case errors.DatabaseError:
			c.log.Error(err.Error())
		default:
			c.log.Warn(err.Error())
		}

		return nil, errors.WrapToJsonRpcError(err)
	}

	return respBuilders.BuildCoordinatesActResponse(status, shipStatus), nil
}
