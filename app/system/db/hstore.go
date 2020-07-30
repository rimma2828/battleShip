package db

import (
	"database/sql/driver"
	"github.com/lib/pq/hstore"
)

type Hstore struct {
	hstore.Hstore
}

func (h Hstore) Value() (driver.Value, error) {
	value, err := h.Hstore.Value()
	if err != nil {
		return nil, err
	}

	if data, ok := value.([]byte); ok {
		return string(data), nil
	} else {
		return value, nil
	}
}
