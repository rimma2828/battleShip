package models

type Ship struct {
	// DB stored fields
	ID     int64 `db:"id"`
	Length int64 `db:"length"`
	Count  int64 `db:"count"`
}
