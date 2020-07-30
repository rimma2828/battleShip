package models

import (
	gen "battleship/app/generated/models"
	appModels "battleship/app/models"
)

func BuildUserModel(params gen.UserCreateParams) (*appModels.User, error) {

	return &appModels.User{
		Name: string(params.UserName),
	}, nil
}
