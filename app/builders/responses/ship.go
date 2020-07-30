package responses

import (
	gen "battleship/app/generated/models"
	"battleship/app/models"
)

func BuildShipResponse(ship *models.Ship) gen.ShipResponse {

	return gen.ShipResponse{
		ID:     ship.ID,
		Length: ship.Length,
		Count:  ship.Count,
	}
}

func BuildShipListResponse(ships []models.Ship) []gen.ShipResponse {
	response := []gen.ShipResponse{}

	for _, ship := range ships {
		response = append(response, BuildShipResponse(&ship))
	}

	return response
}
