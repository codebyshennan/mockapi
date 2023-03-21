package mw

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"bitbucket.org/libertywireless/circles-sandbox/internal"
	"bitbucket.org/libertywireless/circles-sandbox/internal/common/h"
	"bitbucket.org/libertywireless/circles-sandbox/internal/controllers/mockservers"
	"bitbucket.org/libertywireless/circles-sandbox/internal/controllers/sandbox"
	"bitbucket.org/libertywireless/circles-sandbox/internal/controllers/services"
	"bitbucket.org/libertywireless/circles-sandbox/internal/controllers/users"
	"bitbucket.org/libertywireless/circles-sandbox/internal/mw"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var origins = []string{
	"https://sg-kirk.circles.life:7443",
	"http://localhost:3000",
	"https://qsg-sandbox.circles.life",
	"https://psg-sandbox.circles.life",
	"https://qsg-cmsui.circles.life",
	"https://qsg-sandbox-01.circles.life",
	"https://ssg-sandbox.circles.life"}

func RunServer(static embed.FS) {
	// setup dependencies
	provider := internal.NewProvider()
	if provider == nil {

	}

	middlewareProvider := mw.NewMiddlewareProvider(provider)

	// setup outermost router and CORS
	rtr := chi.NewRouter()
	rtr.Use(cors.Handler(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-Request-Id", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	// setup subrouters
	rtr.Use(middleware.Logger)
	rtr.Mount("/api", NewApiRouter(provider, middlewareProvider))
	rtr.Mount("/sandbox", sandbox.NewSandboxHandler(provider, middlewareProvider))
	// rtr.Mount("/sandboxProfile", http.HandlerFunc(sandboxProfileSvc.SandboxProfileHandler))

	fsys := fs.FS(static)
	rtr.Mount("/", http.FileServer(http.FS(fsys)))

	// list all routes
	if err := listRoutes(provider, rtr); err != nil {
		provider.Logger.Errorln(h.InitError, "Logging err: %s\n", err.Error())
	}

	// run server
	if h.FileExists("/etc/server.crt") && h.FileExists("/etc/server.key") {
		provider.Logger.Infoln(h.InitOk, "Serving production server")
		http.ListenAndServeTLS(":3000", "/etc/server.crt", "/etc/server.key", rtr)
	} else {
		provider.Logger.Infoln(h.InitOk, "Serving development server")
		http.ListenAndServe(":3080", rtr)
	}
}

func NewApiRouter(p *internal.Provider, mw *mw.MiddlewareProvider) http.Handler {
	r := chi.NewRouter()

	r.Mount("/v1/auth", users.NewAuthRouter(p, mw))
	r.Mount("/v1/users", users.NewUserRouter(p, mw))
	r.Mount("/v1/services", services.NewServiceRouter(p, mw))
	r.Mount("/v1/swaggers", services.NewSwaggerRouter(p, mw))
	r.Mount("/v1/mockServers", mockservers.NewMockServerRouter(p, mw))

	return r
}

func listRoutes(p *internal.Provider, r chi.Router) error {
	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		s := fmt.Sprintf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		p.Logger.Infoln(h.InitOk, s)
		return nil
	})

	return nil
}
