package responses

import (
	gen "battleship/app/generated/models"
)

func BuildCoordinatesActResponse(status bool, shipStatus string) gen.CoordinatesActResponse {

	return gen.CoordinatesActResponse{
		Status:     status,
		ShipStatus: shipStatus,
	}
}

func BuildCoordinatesResponse(status bool) gen.CoordinatesResponse {

	return gen.CoordinatesResponse{
		Status: status,
	}
}
