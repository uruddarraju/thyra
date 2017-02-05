package runtime

type Object interface {
	GetKind() string
	GetGroup() string
	GetMetadata() map[string]string
}
