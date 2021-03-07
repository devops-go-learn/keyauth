package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	api = &handler{}
)

type handler struct {
	service user.UserServiceClient
	domain  domain.DomainServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	prmaryRouter := router.ResourceRouter("primary_account")
	prmaryRouter.BasePath("users")
	prmaryRouter.Permission(true)
	prmaryRouter.Handle("POST", "/", h.CreatePrimayAccount).AddLabel(label.Create)
	prmaryRouter.Handle("DELETE", "/", h.DestroyPrimaryAccount).AddLabel(label.Delete)

	ramRouter := router.ResourceRouter("ram_account")
	ramRouter.Permission(true)
	ramRouter.BasePath("sub_users")
	ramRouter.Handle("POST", "/", h.CreateSubAccount).AddLabel(label.Create)
	ramRouter.Handle("GET", "/", h.QuerySubAccount).AddLabel(label.List)
	ramRouter.Handle("GET", "/:account", h.DescribeSubAccount).AddLabel(label.Get)
	ramRouter.Handle("PATCH", "/:account", h.PatchSubAccount).AddLabel(label.Update)
	ramRouter.Handle("DELETE", "/:account", h.DestroySubAccount).AddLabel(label.Delete)

	portalRouter := router.ResourceRouter("profile")
	portalRouter.BasePath("profile")
	portalRouter.Handle("GET", "/", h.QueryProfile).AddLabel(label.Get)
	portalRouter.Handle("GET", "/domain", h.QueryDomain).AddLabel(label.Get)
	portalRouter.Handle("PUT", "/", h.PutProfile).AddLabel(label.Update)
	portalRouter.Handle("PATCH", "/", h.PatchProfile).AddLabel(label.Update)

	domRouter := router.ResourceRouter("domain")
	domRouter.BasePath("settings/domain")
	domRouter.Permission(true)
	domRouter.Handle("PUT", "/info", h.UpdateDomainInfo).AddLabel(label.Update)
	domRouter.Handle("PUT", "/security", h.UpdateDomainSecurity).AddLabel(label.Update)

	passRouter := router.ResourceRouter("password")
	passRouter.BasePath("password")
	passRouter.Handle("POST", "/", h.GeneratePassword).AddLabel(label.Create)
	passRouter.Handle("PUT", "/", h.UpdatePassword).AddLabel(label.Update)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.User()
	h.domain = client.Domain()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("user", api)
}
