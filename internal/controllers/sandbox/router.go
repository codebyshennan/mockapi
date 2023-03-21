package sandbox

import (
	"net/http"

	"bitbucket.org/libertywireless/circles-sandbox/internal"

	mw "bitbucket.org/libertywireless/circles-sandbox/internal/mw"
)

func NewSandboxHandler(p *internal.Provider, mw *mw.MiddlewareProvider) http.HandlerFunc {
	controller := NewSandboxController(p)
	return http.HandlerFunc(controller.resolveRoute)
}
