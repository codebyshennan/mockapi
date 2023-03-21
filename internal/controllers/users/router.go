package users

import (
	"github.com/codebyshennan/mockapi/internal"
	mw "github.com/codebyshennan/mockapi/internal/mw"
	"github.com/go-chi/chi/v5"
)

func NewUserRouter(p *internal.Provider, mw *mw.MiddlewareProvider) chi.Router {
	r := chi.NewRouter()
	controller := NewUserController(p)

	r.Use(mw.AuthenticateTokenAndSetContext)
	r.Get("/", controller.getUsers)
	r.Get("/self", controller.getSelf)
	r.Get("/{userId}", controller.getUserById)

	return r
}

func NewAuthRouter(p *internal.Provider, mw *mw.MiddlewareProvider) chi.Router {
	r := chi.NewRouter()
	controller := NewAuthController(p)

	r.Post("/googleLogin", controller.googleLogin)
	r.Post("/logout", controller.logout)

	return r
}
