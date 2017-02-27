package thyra

import (
	"time"
)

// ObjectMeta consists of metadata that is included in every object stored in the api-server.
type ObjectMeta struct {
	// CreatedAt denotes to created timestamp of the object.
	CreatedAt time.Time
	// DeletedAt denotes to deleted timestamp of the object.
	// An object might have this timestamp set without having it cleaned up from the storage
	// to denote a soft delete, which might be followed by an actual delete from the store.
	DeletedAt time.Time
	// UpdateAt timestamp denotes the last update time of the object.
	UpdateAt time.Time
	// Name refers to the name of the object. It is unique amongst the objects of the same kind and api-group.
	Name string
	// UUID refers to a unique identifier for the object in the api-server.
	UUID string
	// Kind refers to the type of the object.
	Kind string
	// Group refers to the API Group the kind of the object belongs to.
	Group string
	// Version represents the version of the object if supported by the storage.
	// +optional
	Version string
	// Labels are a set of key value pairs that can be added to the objects.
	Labels map[string]string
	// Annotations are a set of key value pairs that can be added to the objects.
	Annotations map[string]string
}

// GetName returns the name of the object from ObjectMeta.
func (om *ObjectMeta) GetName() string {
	return om.Name
}

// RestAPIs represents a list of Rest APIs.
type RestAPIs struct {
	ObjectMeta
	Items []*RestAPI
}

// RestAPI refers to an api group specification in the api gateway.
// Inspiration: https://docs.aws.amazon.com/apigateway/api-reference/resource/rest-api/
type RestAPI struct {
	ObjectMeta
	ID          string
	Name        string
	Description string
	CloneFrom   string
	CreatedAt   time.Time
	LastUpdated time.Time
}

// Resources represent a List of Resources in api gateway.
type Resources struct {
	ObjectMeta
	Items []*Resource
}

// Resource represents a type of an object that is to be captured as a rest resource in the api gateway.
type Resource struct {
	ObjectMeta
	ParentResourceId string
	PathPart         string
	Path             string
}

// Methods represent a list of Method objects.
type Methods struct {
	ObjectMeta
	Items []*Method
}

// method represents a HTTP operation and a set of middleware.
// https://docs.aws.amazon.com/apigateway/api-reference/resource/method/
type Method struct {
	ObjectMeta
	Name HttpMethod
	// Authorizers refers to  a list of supported authorizers and also registered custom authorizers,
	// which can be other lambda functions that return a particular response.
	Authorizers       []string
	RequestParameters map[string]string
	RequestModels     []string
}

// TODO: To be defined.
// http://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html
type DomainName struct {
	ObjectMeta
	Name                  string
	CeritifcateName       string
	CertificateBody       string
	CertificatePrivateKey string
	CertificateChain      string
}

// Authorizers represent a list of authorizers.
type Authorizers struct {
	ObjectMeta
	Items []*Authorizer
}

// For custom authorization in future. Not needed to implement now.
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

const (
	IntegrationHTTP       IntegrationType = "HTTP"
	IntegrationHttpProxy  IntegrationType = "HTTP_PROXY"
	IntegrationK8sService IntegrationType = "K8sService"
	IntegrationLambda     IntegrationType = "Lamda"
	IntegrationGRPC       IntegrationType = "GRPC"
	IntegrationMock       IntegrationType = "Mock"
)

type HttpMethod string

const (
	HTTPPost   HttpMethod = "POST"
	HTTPGet    HttpMethod = "GET"
	HTTPPut    HttpMethod = "PUT"
	HTTPDelete HttpMethod = "DELETE"
)

// Integration represents the final destination of the api request handled by the api gateway.
type Integration struct {
	ObjectMeta
	// Type represents the kind of integration.
	Type IntegrationType
	// HttpMethod represents the
	HttpMethod        HttpMethod
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
func (r *RestAPIs) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *RestAPI) GetKind() string         { return "RestApi" }
func (r *RestAPI) GetGroup() string        { return "thyra" }
func (r *RestAPI) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Resources) GetKind() string         { return "Resources" }
func (r *Resources) GetGroup() string        { return "thyra" }
func (r *Resources) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Resource) GetKind() string         { return "Resource" }
func (r *Resource) GetGroup() string        { return "thyra" }
func (r *Resource) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Methods) GetKind() string         { return "Methods" }
func (r *Methods) GetGroup() string        { return "thyra" }
func (r *Methods) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Method) GetKind() string         { return "Method" }
func (r *Method) GetGroup() string        { return "thyra" }
func (r *Method) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Authorizers) GetKind() string         { return "Authorizers" }
func (r *Authorizers) GetGroup() string        { return "thyra" }
func (r *Authorizers) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Authorizer) GetKind() string         { return "Authorizer" }
func (r *Authorizer) GetGroup() string        { return "thyra" }
func (r *Authorizer) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Stages) GetKind() string         { return "Stages" }
func (r *Stages) GetGroup() string        { return "thyra" }
func (r *Stages) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Stage) GetKind() string         { return "Stage" }
func (r *Stage) GetGroup() string        { return "thyra" }
func (r *Stage) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Integrations) GetKind() string         { return "Integrations" }
func (r *Integrations) GetGroup() string        { return "thyra" }
func (r *Integrations) GetMetadata() ObjectMeta { return r.ObjectMeta }

func (r *Integration) GetKind() string         { return "Integration" }
func (r *Integration) GetGroup() string        { return "thyra" }
func (r *Integration) GetMetadata() ObjectMeta { return r.ObjectMeta }
