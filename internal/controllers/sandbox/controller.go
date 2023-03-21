package sandbox

import (
	"net/http"

	"github.com/codebyshennan/mockapi/internal"
	"github.com/codebyshennan/mockapi/internal/common/h"
)

type SandboxController struct {
	provider *internal.Provider
}

func NewSandboxController(p *internal.Provider) (m SandboxController) {
	m = SandboxController{
		provider: p,
	}
	if p == nil {
		p.Logger.Errorln(h.InitError)
	}
	return
}

func (o SandboxController) resolveRoute(w http.ResponseWriter, r *http.Request) {
	sbUri, err := newSandboxUri(r, o.provider)
	if err != nil {
		h.JsonRes(w, http.StatusNotFound, nil, o.provider.Logger)
		return
	}

	sbUri.setResponse(&w)
}
