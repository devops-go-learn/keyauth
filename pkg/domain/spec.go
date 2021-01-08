package domain

import (
	"encoding/json"
	"fmt"

	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/common/types"
	"github.com/infraboard/keyauth/pkg/token/session"
)

// Service is an domain service
type Service interface {
	CreateDomain(ownerID string, req *CreateDomainRequest) (*Domain, error)
	UpdateDomain(req *UpdateDomainInfoRequest) (*Domain, error)
	DescriptionDomain(req *DescribeDomainRequest) (*Domain, error)
	QueryDomain(req *QueryDomainRequest) (*Set, error)
	DeleteDomain(id string) error

	UpdateDomainSecurity(req *UpdateDomainSecurityRequest) (*SecuritySetting, error)
}

// NewQueryDomainRequest 查询domian列表
func NewQueryDomainRequest(page *request.PageRequest) *QueryDomainRequest {
	return &QueryDomainRequest{
		PageRequest:           page,
		Session:               session.NewSession(),
		DescribeDomainRequest: NewDescribeDomainRequest(),
	}
}

// QueryDomainRequest 请求
type QueryDomainRequest struct {
	*session.Session
	*request.PageRequest
	*DescribeDomainRequest
}

// Validate 校验请求合法
func (req *QueryDomainRequest) Validate() error {
	return nil
}

// NewDescribeDomainRequest 查询详情请求
func NewDescribeDomainRequest() *DescribeDomainRequest {
	return &DescribeDomainRequest{}
}

// NewDescribeDomainRequestWithName 查询详情请求
func NewDescribeDomainRequestWithName(name string) *DescribeDomainRequest {
	return &DescribeDomainRequest{
		Name: name,
	}
}

// DescribeDomainRequest 查询domain详情请求
type DescribeDomainRequest struct {
	Name string `json:"name,omitempty"`
}

// Validate todo
func (req *DescribeDomainRequest) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("name required")
	}

	return nil
}

// NewCreateDomainRequest todo
func NewCreateDomainRequest() *CreateDomainRequest {
	return &CreateDomainRequest{
		Enabled: true,
	}
}

// CreateDomainRequest 创建请求
type CreateDomainRequest struct {
	Name           string `bson:"_id" json:"name" validate:"required,lte=60"`               // 公司或者组织名称
	DisplayName    string `bson:"display_name" json:"display_name" validate:"lte=80"`       // 全称
	LogoPath       string `bson:"logo_path" json:"logo_path" validate:"lte=200"`            // 公司LOGO图片的URL
	Description    string `bson:"description" json:"description" validate:"lte=400"`        // 描述
	Enabled        bool   `bson:"enabled" json:"enabled"`                                   // 域状态, 是否需要冻结该域, 冻结时, 该域下面所有用户禁止登录
	Size           string `bson:"size" json:"size" validate:"lte=20"`                       // 规模: 50人以下, 50~100, ...
	Location       string `bson:"location" json:"location" validate:"lte=20"`               // 位置: 指城市, 比如 中国,四川,成都
	Address        string `bson:"address" json:"address" validate:"lte=200"`                // 地址: 比如环球中心 10F 1034
	Industry       string `bson:"industry" json:"industry" validate:"lte=100"`              // 所属行业: 比如, 互联网
	Fax            string `bson:"fax" json:"fax" validate:"lte=40"`                         // 传真:
	Phone          string `bson:"phone" json:"phone" validate:"lte=20"`                     // 固话:
	ContactsName   string `bson:"contacts_name" json:"contacts_name" validate:"lte=30"`     // 联系人姓名
	ContactsTitle  string `bson:"contacts_title" json:"contacts_title" validate:"lte=40"`   // 联系人职位
	ContactsMobile string `bson:"contacts_mobile" json:"contacts_mobile" validate:"lte=20"` // 联系人电话
	ContactsEmail  string `bson:"contacts_email" json:"contacts_email" validate:"lte=40"`   // 联系人邮箱
}

// Validate 校验请求是否合法
func (req *CreateDomainRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateDomainRequest) String() string {
	return fmt.Sprint(*req)
}

// Patch todo
func (req *CreateDomainRequest) Patch(data *CreateDomainRequest) {
	patchData, _ := json.Marshal(data)
	json.Unmarshal(patchData, req)
}

// NewPutDomainRequest todo
func NewPutDomainRequest() *UpdateDomainInfoRequest {
	return &UpdateDomainInfoRequest{
		UpdateMode:          types.PutUpdateMode,
		CreateDomainRequest: NewCreateDomainRequest(),
	}
}

// NewPatchDomainRequest todo
func NewPatchDomainRequest() *UpdateDomainInfoRequest {
	return &UpdateDomainInfoRequest{
		UpdateMode:          types.PatchUpdateMode,
		CreateDomainRequest: NewCreateDomainRequest(),
	}
}

// NewPutDomainSecurityRequest todo
func NewPutDomainSecurityRequest() *UpdateDomainSecurityRequest {
	return &UpdateDomainSecurityRequest{
		UpdateMode:      types.PutUpdateMode,
		SecuritySetting: NewDefaultSecuritySetting(),
	}
}

// UpdateDomainSecurityRequest todo
type UpdateDomainSecurityRequest struct {
	Name       string           `json:"domain_name"`
	UpdateMode types.UpdateMode `json:"update_mode"`
	*SecuritySetting
}

// Validate todo
func (req *UpdateDomainSecurityRequest) Validate() error {
	return validate.Struct(req)
}

// UpdateDomainInfoRequest todo
type UpdateDomainInfoRequest struct {
	UpdateMode types.UpdateMode `json:"update_mode"`
	*CreateDomainRequest
}

// Validate 更新校验
func (req *UpdateDomainInfoRequest) Validate() error {
	if req.CreateDomainRequest == nil {
		return fmt.Errorf("domain base info required")
	}

	if err := req.CreateDomainRequest.Validate(); err != nil {
		return err
	}

	return nil
}
