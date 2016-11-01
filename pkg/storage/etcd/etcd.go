package etcd

type Storage interface {
	Name() string
}

type EtcdStorage struct{}

func NewEtcdStorage() Storage {
	return &EtcdStorage{}
}

func (es *EtcdStorage) Name() string {
	return "Etcd"
}
