package runtime

type Object interface {
	GetKind() string
	GetGroup() string
	GetName() string
	GetMetadata() map[string]string
}
