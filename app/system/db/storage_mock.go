package db

type StorageMock struct {
}

func NewStorageMock() StorageMock {
	return StorageMock{}
}

func (i StorageMock) BeginTx() (StorageTx, error) {
	return StorageTxMock{}, nil
}

func (i StorageMock) Exec(query string, params ...interface{}) error {
	return nil
}

func (i StorageMock) Query(coll interface{}, query string, params ...interface{}) error {
	return nil
}

type StorageTxMock struct {
}

func (i StorageTxMock) BeginTx() (StorageTx, error) {
	return i, nil
}

func (i StorageTxMock) Exec(query string, params ...interface{}) error {
	return nil
}

func (i StorageTxMock) Query(coll interface{}, query string, params ...interface{}) error {
	return nil
}

func (i StorageTxMock) Commit() error {
	return nil
}

func (i StorageTxMock) Rollback() error {
	return nil
}
