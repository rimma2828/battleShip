package user

import (
	"context"

	modelsBuilders "battleship/app/builders/models"
	respBuilders "battleship/app/builders/responses"
	"battleship/app/errors"
	"battleship/app/generated/models"
)

func (c *User) Create(ctx context.Context, params models.UserCreateParams) (interface{}, error) {
	userModel, _ := modelsBuilders.BuildUserModel(params)
	user, err := c.processor.Create(userModel)
	if err != nil {
		switch err.(type) {
		case errors.DatabaseError:
			c.log.Error(err.Error())
		default:
			c.log.Warn(err.Error())
		}
		return nil, errors.WrapToJsonRpcError(err)
	}
	return respBuilders.BuildUserResponse(user), nil
}
