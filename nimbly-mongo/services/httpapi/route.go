package httpapi

import (
	"chi-rest/bootstrap"
	"chi-rest/services/httpapi/handler"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// RegisterRoutes all routes for the apps
func RegisterRoutes(r *chi.Mux, app *bootstrap.App) {

	//The url pointing to API definition"
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(app.Config.GetString("app.app_host")+"/swagger/doc.json"),
	))

	h := handler.Contract{app}
	r.Route("/v1", func(r chi.Router) {
		r.Get("/", h.Hello)
	})
}
