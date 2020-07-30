package models

type User struct {
	// DB stored fields
	ID   int64  `db:"id"`
	Name string `db:"username"`
}
