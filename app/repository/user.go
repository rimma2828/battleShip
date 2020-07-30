package repository

import (
	"fmt"

	"battleship/app/errors"
	appModels "battleship/app/models"
	"battleship/app/system/db"
)

type UserRepository interface {
	Get(client db.Storage, id int64) (*appModels.User, error)
	GetByName(client db.Storage, name string) (*appModels.User, error)
	Save(client db.Storage, user *appModels.User) (int64, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (c *userRepository) Get(client db.Storage, id int64) (*appModels.User, error) {
	var user []appModels.User

	err := client.Query(&user,
		`select 
		 id, 
		 username
		from users where id = $1
		limit 1;`,
		id)

	if err != nil {
		return nil, errors.NewDatabaseError(errors.DatabaseBasicError, err)
	}
	if len(user) == 0 {
		return nil, fmt.Errorf("пользователь не найден")
	}

	return &user[0], nil
}

func (c *userRepository) GetByName(client db.Storage, name string) (*appModels.User, error) {
	var user []appModels.User

	err := client.Query(&user,
		`select 
		 id, 
		 username
		from users where username = $1
		limit 1;`,
		name)

	if err != nil {
		return nil, errors.NewDatabaseError(errors.DatabaseBasicError, err)
	}
	if user != nil {
		return &user[0], nil
	} else {
		return nil, nil
	}

}

func (c *userRepository) Save(client db.Storage, user *appModels.User) (int64, error) {
	var sId []int64
	q := `insert into users (username) VALUES ($1) returning id;`
	err := client.Query(&sId,
		q,
		user.Name,
	)
	if err != nil {
		return 0, errors.NewDatabaseError(errors.DatabaseBasicError, err)
	}

	return sId[0], nil
}
