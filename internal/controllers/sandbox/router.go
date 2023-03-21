package sandbox

import (
	"net/http"

	"github.com/codebyshennan/mockapi/internal"

	mw "github.com/codebyshennan/mockapi/internal/mw"
)

func NewSandboxHandler(p *internal.Provider, mw *mw.MiddlewareProvider) http.HandlerFunc {
	controller := NewSandboxController(p)
	return http.HandlerFunc(controller.resolveRoute)
}
