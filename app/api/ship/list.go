package ship

import (
	"context"

	builders "battleship/app/builders/responses"
	"battleship/app/errors"
	"battleship/app/generated/models"
)

func (c *Ship) List(ctx context.Context, params models.ShipListParams) (interface{}, error) {
	ship, err := c.processor.GetList()
	if err != nil {
		return nil, errors.WrapToJsonRpcError(err)
	}

	return builders.BuildShipListResponse(ship), nil
}
