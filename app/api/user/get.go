package user

import (
	"context"

	"battleship/app/errors"
	"battleship/app/generated/models"
)

func (c *User) Get(ctx context.Context, params models.UserGetParams) (interface{}, error) {
	user, err := c.processor.Get(params.ID)
	if err != nil {
		switch err.(type) {
		case errors.DatabaseError:
			c.log.Error(err.Error())
		default:
			c.log.Warn(err.Error())
		}
		return nil, errors.WrapToJsonRpcError(err)
	}
	return models.UserGetResponse{ID: user.ID, UserName: user.Name}, nil
}
