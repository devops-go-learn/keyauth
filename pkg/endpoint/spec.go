package endpoint

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/types/ftime"

	"github.com/infraboard/keyauth/pkg/token/session"
	"github.com/infraboard/keyauth/pkg/user/types"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Service token管理服务
type Service interface {
	DescribeEndpoint(req *DescribeEndpointRequest) (*Endpoint, error)
	QueryEndpoints(req *QueryEndpointRequest) (*Set, error)
	Registry(req *RegistryRequest) error
	DeleteEndpoint(req *DeleteEndpointRequest) error
}

// NewRegistryRequest 注册请求
func NewRegistryRequest(version string, entries []*router.Entry) *RegistryRequest {
	return &RegistryRequest{
		Version: version,
		Entries: entries,
		Session: session.NewSession(),
	}
}

// NewDefaultRegistryRequest todo
func NewDefaultRegistryRequest() *RegistryRequest {
	return &RegistryRequest{
		Session: session.NewSession(),
		Entries: []*router.Entry{},
	}
}

// RegistryRequest 服务注册请求
type RegistryRequest struct {
	*session.Session
	Version string          `json:"version" validate:"required,lte=32"`
	Entries []*router.Entry `json:"entries"`
}

// Validate 校验注册请求合法性
func (req *RegistryRequest) Validate() error {
	if len(req.Entries) == 0 {
		return fmt.Errorf("must require *router.Entry")
	}

	tk := req.GetToken()
	if tk == nil {
		return fmt.Errorf("token required when service endpoints registry")
	}

	if !tk.UserType.Is(types.UserType_SERVICE) {
		return fmt.Errorf("only service account can registry endpoints")
	}

	return validate.Struct(req)
}

// Endpoints 功能列表
func (req *RegistryRequest) Endpoints(serviceID string) []*Endpoint {
	eps := make([]*Endpoint, 0, len(req.Entries))
	for i := range req.Entries {
		ep := &Endpoint{
			ID:        GenHashID(serviceID, req.Entries[i].Path, req.Entries[i].Method),
			CreateAt:  ftime.Now(),
			UpdateAt:  ftime.Now(),
			ServiceID: serviceID,
			Version:   req.Version,
			Entry:     *req.Entries[i],
		}
		eps = append(eps, ep)
	}
	return eps
}

// NewQueryEndpointRequestFromHTTP 列表查询请求
func NewQueryEndpointRequestFromHTTP(r *http.Request) *QueryEndpointRequest {
	page := request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()

	return &QueryEndpointRequest{
		PageRequest:  page,
		ServiceID:    qs.Get("service_id"),
		Path:         qs.Get("path"),
		Method:       qs.Get("method"),
		FunctionName: qs.Get("function_name"),
		Resource:     qs.Get("resource"),
	}
}

// NewQueryEndpointRequest 列表查询请求
func NewQueryEndpointRequest(pageReq *request.PageRequest) *QueryEndpointRequest {
	return &QueryEndpointRequest{
		PageRequest: pageReq,
	}
}

// QueryEndpointRequest 查询应用列表
type QueryEndpointRequest struct {
	*request.PageRequest
	ServiceID    string
	Path         string
	Method       string
	FunctionName string
	Resource     string
	Labels       map[string]string
}

// NewDescribeEndpointRequestWithID todo
func NewDescribeEndpointRequestWithID(id string) *DescribeEndpointRequest {
	return &DescribeEndpointRequest{ID: id}
}

// DescribeEndpointRequest todo
type DescribeEndpointRequest struct {
	ID string
}

// Validate 校验
func (req *DescribeEndpointRequest) Validate() error {
	if req.ID == "" {
		return fmt.Errorf("endpoint id is required")
	}

	return nil
}

// NewDeleteEndpointRequestWithServiceID todo
func NewDeleteEndpointRequestWithServiceID(id string) *DeleteEndpointRequest {
	return &DeleteEndpointRequest{ServiceID: id}
}

// DeleteEndpointRequest todo
type DeleteEndpointRequest struct {
	ServiceID string
}
