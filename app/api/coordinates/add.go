package coordinates

import (
	"context"

	respBuilders "battleship/app/builders/responses"
	"battleship/app/errors"
	gen "battleship/app/generated/models"
)

func (c *Coordinates) Add(ctx context.Context, params gen.CoordinatesAddParams) (interface{}, error) {
	status, err := c.processor.Add(&params)
	if err != nil {
		switch err.(type) {
		case errors.DatabaseError:
			c.log.Error(err.Error())
		default:
			c.log.Warn(err.Error())
		}

		return nil, errors.WrapToJsonRpcError(err)
	}
	return respBuilders.BuildCoordinatesResponse(status), nil
}
