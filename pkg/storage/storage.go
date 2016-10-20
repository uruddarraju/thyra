package storage

type Storage interface {
	Name() string
}

type EtcdStorage struct{}

func NewDefaultStorage() Storage {
	return &EtcdStorage{}
}

func (es *EtcdStorage) Name() string {
	return "Etcd"
}
