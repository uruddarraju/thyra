package api

import (
	"time"
)

var (
	HTTPIntegration      = "http"
	HTTPProxyIntegration = "http_proxy"
	LambdaProxy          = "lambda"
)

type ObjectMeta struct {
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
	Name      string
	UUID      string
	Kind 	  string
	Group 	  string
	Version   string
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

func(r *RestAPIs) GetKind() { return "RestApis" }
func(r *RestAPIs) GetGroup() { return "thyra" }

func(r *RestAPI) GetKind() { return "RestApi" }
func(r *RestAPI) GetGroup() { return "thyra" }

func(r *Resources) GetKind() { return "Resources" }
func(r *Resources) GetGroup() { return "thyra" }

func(r *Resource) GetKind() { return "Resource" }
func(r *Resource) GetGroup() { return "thyra" }

func(r *Methods) GetKind() { return "Methods" }
func(r *Methods) GetGroup() { return "thyra" }

func(r *Method) GetKind() { return "Method" }
func(r *Method) GetGroup() { return "thyra" }

func(r *Authorizers) GetKind() { return "Authorizers" }
func(r *Authorizers) GetGroup() { return "thyra" }

func(r *Authorizer) GetKind() { return "Authorizer" }
func(r *Authorizer) GetGroup() { return "thyra" }

func(r *Stages) GetKind() { return "Stages" }
func(r *Stages) GetGroup() { return "thyra" }

func(r *Stage) GetKind() { return "Stage" }
func(r *Stage) GetGroup() { return "thyra" }

func(r *Integrations) GetKind() { return "Integrations" }
func(r *Integrations) GetGroup() { return "thyra" }

func(r *Integration) GetKind() { return "Integration" }
func(r *Integration) GetGroup() { return "thyra" }


