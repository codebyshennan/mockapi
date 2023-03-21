package mockservers

import (
	"bitbucket.org/libertywireless/circles-sandbox/internal"
	mw "bitbucket.org/libertywireless/circles-sandbox/internal/mw"
	"github.com/go-chi/chi/v5"
)

func NewMockServerRouter(p *internal.Provider, mw *mw.MiddlewareProvider) chi.Router {
	r := chi.NewRouter()
	controller := NewMockServerController(p)

	// r.Use(mw.AuthenticateTokenAndSetUserId)

	r.Use(mw.AuthenticateTokenAndSetContext)
	r.Post("/", controller.createMockServer)
	r.Get("/", controller.getMockServers)
	r.Get("/{mockServerId}", controller.getMockServerById)
	r.Patch("/{mockServerId}", controller.updateMockServer)

	return r
}
