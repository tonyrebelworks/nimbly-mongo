package journeyplan

import (
	"chi-rest/bootstrap"
	"chi-rest/services/journeyplan/handler"

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
		// r.Get("/", h.Hello)

		//Journey CMS
		r.Get("/journey", h.GetAllJourney)
		r.Get("/journey/{code}", h.GetDetailJourney)
		r.Post("/journey", h.AddJourney)
		r.Put("/journey/{code}", h.UpdateJourney)
		r.Delete("/journey/{code}", h.DeleteJourney)
		r.Get("/journey/report/{code}", h.GetReportJourney)

		//Journey Mobile
		r.Get("/journeymobile", h.GetAllJourneyMobile)
		r.Get("/journeymobile/{code}", h.GetDetailJourneyMobile)
		r.Put("/journey/time", h.UpdateTimeJourney)
		r.Post("/journey/trackingtime", h.AddTrackingTimeJourney)
		r.Post("/journey/url", h.AddURLFirebase)

	})
}
