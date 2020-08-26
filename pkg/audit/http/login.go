package http

import (
	"net/http"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/audit"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) QueryLoginLog(w http.ResponseWriter, r *http.Request) {
	tk, err := pkg.GetTokenFromContext(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req, err := audit.NewQueryLoginRecordRequestFromHTTP(r)
	if err != nil {
		response.Failed(w, exception.NewBadRequest("validate request error, %s", err))
		return
	}
	req.WithToken(tk)

	set, err := h.service.QueryLoginRecord(req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
	return
}
