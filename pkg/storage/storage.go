package storage

type Storage interface{}

type EtcdStorage struct{}

func NewDefaultStorage() *Storage {
	return &EtcdStorage{}
}
