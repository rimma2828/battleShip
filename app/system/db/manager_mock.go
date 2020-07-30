package db

type storageManagerMock struct {
	rw Storage
	ro Storage
}

func NewStorageManagerMock() StorageManager {
	return &storageManager{
		rw: NewStorageMock(),
		ro: NewStorageMock(),
	}
}

func (s *storageManagerMock) GetRW() Storage {
	return s.rw
}

func (s *storageManagerMock) GetRO() Storage {
	return s.ro
}
