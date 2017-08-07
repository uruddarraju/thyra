package thyra

type GroupVersionKind struct {
	Group   APIGroup
	Version string
	Kind    string
}

type APIGroup struct{}
