package policy

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/token/session"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Service 策略服务
type Service interface {
	CreatePolicy(*CreatePolicyRequest) (*Policy, error)
	QueryPolicy(*QueryPolicyRequest) (*Set, error)
	DescribePolicy(*DescribePolicyRequest) (*Policy, error)
	DeletePolicy(*DeletePolicyRequest) error
}

// NewQueryPolicyRequestFromHTTP 列表查询请求
func NewQueryPolicyRequestFromHTTP(r *http.Request) *QueryPolicyRequest {
	page := request.NewPageRequestFromHTTP(r)
	req := NewQueryPolicyRequest(page)

	qs := r.URL.Query()
	req.Account = qs.Get("account")
	req.RoleID = qs.Get("role_id")
	req.NamespaceID = qs.Get("namespace_id")
	req.WithRole = qs.Get("with_role") == "true"
	req.WithNamespace = qs.Get("with_namespace") == "true"
	return req
}

// NewQueryPolicyRequest 列表查询请求
func NewQueryPolicyRequest(pageReq *request.PageRequest) *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Session:       session.NewSession(),
		PageRequest:   pageReq,
		WithRole:      false,
		WithNamespace: false,
	}
}

// QueryPolicyRequest 获取子账号列表
type QueryPolicyRequest struct {
	*request.PageRequest
	*session.Session

	Account       string `json:"account,omitempty"`
	RoleID        string `json:"role_id,omitempty"`
	NamespaceID   string `json:"namespace_id,omitempty"`
	Type          *Type  `json:"type,omitempty"`
	WithRole      bool   `json:"with_role,omitempty"`
	WithNamespace bool   `json:"with_namespace,omitempty"`
}

// Validate 校验请求是否合法
func (req *QueryPolicyRequest) Validate() error {
	return validate.Struct(req)
}

// NewDescriptPolicyRequest new实例
func NewDescriptPolicyRequest() *DescribePolicyRequest {
	return &DescribePolicyRequest{
		Session: session.NewSession(),
	}
}

// DescribePolicyRequest todo
type DescribePolicyRequest struct {
	*session.Session
	ID string
}

// Validate todo
func (req *DescribePolicyRequest) Validate() error {
	if req.ID == "" {
		return fmt.Errorf("policy id required")
	}

	return nil
}

// NewDeletePolicyRequestWithID todo
func NewDeletePolicyRequestWithID(id string) *DeletePolicyRequest {
	req := NewDeletePolicyRequest()
	req.ID = id
	return req
}

// NewDeletePolicyRequestWithNamespaceID todo
func NewDeletePolicyRequestWithNamespaceID(namespaceID string) *DeletePolicyRequest {
	req := NewDeletePolicyRequest()
	req.NamespaceID = namespaceID
	return req
}

// NewDeletePolicyRequestWithAccount todo
func NewDeletePolicyRequestWithAccount(account string) *DeletePolicyRequest {
	req := NewDeletePolicyRequest()
	req.Account = account
	return req
}

// NewDeletePolicyRequestWithRoleID todo
func NewDeletePolicyRequestWithRoleID(roleID string) *DeletePolicyRequest {
	req := NewDeletePolicyRequest()
	req.RoleID = roleID
	return req
}

// NewDeletePolicyRequest todo
func NewDeletePolicyRequest() *DeletePolicyRequest {
	return &DeletePolicyRequest{
		Session: session.NewSession(),
	}
}

// DeletePolicyRequest todo
type DeletePolicyRequest struct {
	*session.Session `json:"-"`
	ID               string `json:"id,omitempty"`
	Account          string `json:"account,omitempty"`
	RoleID           string `json:"role_id,omitempty"`
	NamespaceID      string `json:"namespace_id,omitempty"`
	Type             *Type  `json:"type,omitempty"`
}

// Validate todo
func (req *DeletePolicyRequest) Validate() error {
	tk := req.GetToken()
	if tk == nil {
		return fmt.Errorf("token required")
	}

	if req.ID == "" {
		return fmt.Errorf("policy id required")
	}

	return nil
}
