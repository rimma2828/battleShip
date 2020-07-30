package db

import "time"

type Storage interface {
	Exec(query string, params ...interface{}) error
	Query(coll interface{}, query string, params ...interface{}) error
	BeginTx() (StorageTx, error)
}

type StorageTx interface {
	Storage
	Commit() error
	Rollback() error
}

type StorageManager interface {
	GetRW() Storage
	GetRO() Storage
}

type storageManagerConfig interface {
	GetRWDSN() string
	GetRODSN() string
	GetConnectionsMaxIdle() int
	GetConnectionsMaxOpen() int
	GetConnectionsMaxLifetime() time.Duration
}

type storageManager struct {
	rw Storage
	ro Storage
}

func NewStorageManager(c storageManagerConfig) (StorageManager, error) {
	var err error
	res := &storageManager{}
	res.rw, err = NewPQDatabase(c.GetRWDSN(), c.GetConnectionsMaxIdle(), c.GetConnectionsMaxOpen(), c.GetConnectionsMaxLifetime())
	if err != nil {
		return nil, err
	}

	res.ro, err = NewPQDatabase(c.GetRODSN(), c.GetConnectionsMaxIdle(), c.GetConnectionsMaxOpen(), c.GetConnectionsMaxLifetime())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *storageManager) GetRW() Storage {
	return s.rw
}

func (s *storageManager) GetRO() Storage {
	return s.ro
}
