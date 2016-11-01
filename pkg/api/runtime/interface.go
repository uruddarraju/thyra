package api

type Object interface {
	GetKind() string
	GetGroup() string
}
