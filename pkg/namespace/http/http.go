package http

import (
	"errors"

	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/namespace"
)

var (
	api = &handler{}
)

type handler struct {
	service namespace.NamespaceServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("namespace")
	// 获取自己的namespace 不需做权限限制
	r.BasePath("self/namespaces")
	r.Handle("GET", "/", h.ListSelfNamespace).AddLabel(label.List)

	// 需要做权限限制
	r.BasePath("namespaces")
	r.Permission(true)
	r.Handle("POST", "/", h.Create).AddLabel(label.Create)
	r.Handle("GET", "/", h.List).AddLabel(label.List)
	r.Handle("GET", "/:id", h.Get).AddLabel(label.Get)
	r.Handle("DELETE", "/:id", h.Delete).AddLabel(label.Delete)

}

func (h *handler) Config() error {
	if pkg.Namespace == nil {
		return errors.New("denpence namespace service is nil")
	}

	h.service = pkg.Namespace
	return nil
}

func init() {
	pkg.RegistryHTTPV1("namespace", api)
}
