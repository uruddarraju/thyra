package etcd

// TODO: Implement an etcd storage

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
