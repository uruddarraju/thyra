package api

type GroupVersionKind struct {
	Group   APIGroup
	Version string
	Kind    string
}

type APIGroup struct {
}

type RestAPIs struct {
	Items []RestAPI
}

type RestAPI struct {
	Name        string
	Description string
	CloneFrom   string
}

// http://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html
type DomainName struct {
	Name                  string
	CeritifcateName       string
	CertificateBody       string
	CertificatePrivateKey string
	CertificateChain      string
}
