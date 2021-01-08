package application

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/keyauth/pkg/token/session"
	"github.com/infraboard/mcube/http/request"
)

const (
	// DefaultAccessTokenExpireSecond token默认过期时长
	DefaultAccessTokenExpireSecond = 3600
	// DefaultRefreshTokenExpiredSecond 刷新token默认过期时间
	DefaultRefreshTokenExpiredSecond = DefaultAccessTokenExpireSecond * 4
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// Service 用户服务
type Service interface {
	UserInterface
	AdminInterface
}

// UserInterface todo
type UserInterface interface {
	CreateUserApplication(req *CreateApplicatonRequest) (*Application, error)
	DeleteApplication(id string) error
	DescriptionApplication(req *DescriptApplicationRequest) (*Application, error)
	QueryApplication(req *QueryApplicationRequest) (*Set, error)
}

// AdminInterface todo
type AdminInterface interface {
	CreateBuildInApplication(req *CreateApplicatonRequest) (*Application, error)
	GetBuildInApplication(name string) (*Application, error)
}

// NewDescriptApplicationRequest new实例
func NewDescriptApplicationRequest() *DescriptApplicationRequest {
	return &DescriptApplicationRequest{
		Session: session.NewSession(),
	}
}

// DescriptApplicationRequest 查询应用详情
type DescriptApplicationRequest struct {
	*session.Session
	ID       string `json:"id,omitempty"`
	ClientID string `json:"client_id,omitempty"`
}

// Validate 校验详情查询请求
func (req *DescriptApplicationRequest) Validate() error {
	if req.ID == "" && req.ClientID == "" {
		return errors.New("id or client_id is required")
	}

	return nil
}

// NewQueryApplicationRequest 列表查询请求
func NewQueryApplicationRequest(pageReq *request.PageRequest) *QueryApplicationRequest {
	return &QueryApplicationRequest{
		PageRequest: pageReq,
	}
}

// QueryApplicationRequest 查询应用列表
type QueryApplicationRequest struct {
	*request.PageRequest
	UserID  string
	Account string
}

// NewCreateApplicatonRequest 请求
func NewCreateApplicatonRequest() *CreateApplicatonRequest {
	return &CreateApplicatonRequest{
		Session:                   session.NewSession(),
		AccessTokenExpireSecond:   DefaultAccessTokenExpireSecond,
		RefreshTokenExpiredSecond: DefaultRefreshTokenExpiredSecond,
		ClientType:                Public,
	}
}

// CreateApplicatonRequest 创建应用请求
type CreateApplicatonRequest struct {
	*session.Session          `bson:"-" json:"-"`
	Name                      string     `bson:"name" json:"name,omitempty" validate:"required,lte=30"`          // 应用名称
	Website                   string     `bson:"website" json:"website,omitempty" validate:"lte=200"`            // 应用的网站地址
	LogoImage                 string     `bson:"logo_image" json:"logo_image,omitempty" validate:"lte=200"`      // 应用的LOGO
	Description               string     `bson:"description" json:"description,omitempty" validate:"lte=1000"`   // 应用简单的描述
	RedirectURI               string     `bson:"redirect_uri" json:"redirect_uri,omitempty" validate:"lte=200"`  // 应用重定向URI, Oauht2时需要该参数
	AccessTokenExpireSecond   int64      `bson:"access_token_expire_second" json:"access_token_expire_second"`   // 应用申请的token的过期时间
	RefreshTokenExpiredSecond int64      `bson:"refresh_token_expire_second" json:"refresh_token_expire_second"` // 刷新token过期时间
	ClientType                ClientType `bson:"client_type" json:"client_type,omitempty"`                       // 客户端类型
}

// Validate 请求校验
func (req *CreateApplicatonRequest) Validate() error {
	return validate.Struct(req)
}
