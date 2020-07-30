package responses

import (
	gen "battleship/app/generated/models"
	"battleship/app/models"
)

func BuildUserResponse(user *models.User) gen.UserResponse {

	return gen.UserResponse{
		ID:       user.ID,
		UserName: user.Name,
	}
}
