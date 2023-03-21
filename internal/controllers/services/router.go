package services

import (
	"github.com/codebyshennan/mockapi/internal"
	mw "github.com/codebyshennan/mockapi/internal/mw"
	"github.com/go-chi/chi/v5"
)

func NewServiceRouter(p *internal.Provider, mw *mw.MiddlewareProvider) chi.Router {
	r := chi.NewRouter()
	controller := NewServiceController(p)

	r.Use(mw.AuthenticateTokenAndSetContext)
	r.Post("/", controller.createService)
	r.Get("/", controller.getServices)
	r.Get("/{serviceId}", controller.getServiceById)

	return r
}

func NewSwaggerRouter(p *internal.Provider, mw *mw.MiddlewareProvider) chi.Router {
	r := chi.NewRouter()
	controller := NewSwaggerController(p)

	r.Use(mw.AuthenticateTokenAndSetContext)
	r.Post("/", controller.createSwagger)
	r.Get("/", controller.getSwaggers)
	r.Get("/{swaggerId}", controller.getSwaggerById)

	return r
}
