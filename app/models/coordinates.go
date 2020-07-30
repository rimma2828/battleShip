package models

type Coordinates struct {
	// DB stored fields
	UserId int64 `db:"user_id"`
	ShipId int64 `db:"ship_id"`
	XCoord int64 `db:"x_coord"`
	YCoord int64 `db:"y_coord"`
}
