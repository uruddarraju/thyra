package api

import (
	"time"
)

var (
	HTTPIntegration      = "http"
	HTTPProxyIntegration = "http_proxy"
	LambdaProxy          = "lambda"
	GRPCProxy            = "grpc"
)

type ObjectMeta struct {
	CreatedAt   time.Time
	DeletedAt   time.Time
	UpdateAt    time.Time
	Name        string
	UUID        string
	Kind        string
	Group       string
	Version     string
	Labels      map[string]string
	Annotations map[string]string
}

func (om *ObjectMeta) GetName() {
	return om.Name
}

type RestAPIs struct {
	ObjectMeta
	Items []*RestAPI
}

type RestAPI struct {
	ObjectMeta
	ID          string
	Name        string
	Description string
	CloneFrom   string
	CreatedAt   time.Time
	LastUpdated time.Time
}

type Resources struct {
	ObjectMeta
	Items []*Resource
}

type Resource struct {
	ObjectMeta
	ParentResourceId string
	PathPart         string
	Path             string
}

type Methods struct {
	ObjectMeta
	Items []*Method
}

type HttpMethod string

// https://docs.aws.amazon.com/apigateway/api-reference/resource/method/
type Method struct {
	ObjectMeta
	Name HttpMethod
	// Can be a list of supported authorizers and also registered custom authorizers, which can be other lambda functions that return a particular response.
	Authorizers       []string
	RequestParameters map[string]string
	RequestModels     []string
}

// http://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html
type DomainName struct {
	ObjectMeta
	Name                  string
	CeritifcateName       string
	CertificateBody       string
	CertificatePrivateKey string
	CertificateChain      string
}

type Authorizers struct {
	ObjectMeta
	Items []*Authorizer
}

// For custom authorization in future. Not needed to implement now
type Authorizer struct {
	ObjectMeta
	Name                       string
	AuthorizerURI              string
	AuthorizerCredentials      string
	AuthorizerResultTTLSeconds int
}

type BasePathMapping struct {
	ObjectMeta
	BasePath    string
	RestAPIName string
	StageName   string
}

// https://docs.aws.amazon.com/apigateway/api-reference/resource/stages/
// A stage corresponds to a version of the API in service
// In each stage, you can configure stage-level throttling settings, in addition to enabling or disabling API cache or CloudWatch logs for the API's requests and responses.
type Stage struct {
	ObjectMeta
	Name           string
	Description    string
	CacheSpec      CacheSpec
	Deployment     string
	MethodSettings []*MethodSetting
	CreatedAt      time.Time
	LastUpdatedAt  time.Time
}

type MethodSetting struct {
	// TODO: This should be "resource/method" so that we can do regex on the resources
	Method string

	// Would be authorizers that would be an override on top of the api definition
	Authorizers []string

	// Would be middlewares like rate limiting, throttling that would be an override on top of the api definition
	Middlewares []string
}

type CacheSpec struct {
	Enabled         bool
	CacheSize       string
	CacheTTL        int
	CacheEncryption bool
}

type Stages struct {
	ObjectMeta
	Items []*Stage
}

type IntegrationType string

type Integration struct {
	ObjectMeta
	Type              IntegrationType
	HttpMethod        string
	URI               string
	Credentials       string
	RequestParameters map[string]string
}

type Integrations struct {
	ObjectMeta
	Items []*Integration
}

func (r *RestAPIs) GetKind() string         { return "RestApis" }
func (r *RestAPIs) GetGroup() string        { return "thyra" }
func (r *RestAPIs) GetMetadata() ObjectMeta { return nil }

func (r *RestAPI) GetKind() string         { return "RestApi" }
func (r *RestAPI) GetGroup() string        { return "thyra" }
func (r *RestAPI) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Resources) GetKind() string         { return "Resources" }
func (r *Resources) GetGroup() string        { return "thyra" }
func (r *Resources) GetMetadata() ObjectMeta { return nil }

func (r *Resource) GetKind() string         { return "Resource" }
func (r *Resource) GetGroup() string        { return "thyra" }
func (r *Resource) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Methods) GetKind() string         { return "Methods" }
func (r *Methods) GetGroup() string        { return "thyra" }
func (r *Methods) GetMetadata() ObjectMeta { return nil }

func (r *Method) GetKind() string         { return "Method" }
func (r *Method) GetGroup() string        { return "thyra" }
func (r *Method) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Authorizers) GetKind() string         { return "Authorizers" }
func (r *Authorizers) GetGroup() string        { return "thyra" }
func (r *Authorizers) GetMetadata() ObjectMeta { return nil }

func (r *Authorizer) GetKind() string         { return "Authorizer" }
func (r *Authorizer) GetGroup() string        { return "thyra" }
func (r *Authorizer) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Stages) GetKind() string         { return "Stages" }
func (r *Stages) GetGroup() string        { return "thyra" }
func (r *Stages) GetMetadata() ObjectMeta { return nil }

func (r *Stage) GetKind() string         { return "Stage" }
func (r *Stage) GetGroup() string        { return "thyra" }
func (r *Stage) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Integrations) GetKind() string         { return "Integrations" }
func (r *Integrations) GetGroup() string        { return "thyra" }
func (r *Integrations) GetMetadata() ObjectMeta { return nil }

func (r *Integration) GetKind() string         { return "Integration" }
func (r *Integration) GetMetadata() ObjectMeta { return r.ObjectMeta }
func (r *Integration) GetGroup() string        { return "thyra" }
