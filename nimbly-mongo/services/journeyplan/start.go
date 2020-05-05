package journeyplan

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"chi-rest/bootstrap"
	_ "chi-rest/services/journeyplan/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/valve"
	"github.com/urfave/cli/v2"
)

// API ...
type API struct {
	*bootstrap.App
}

// Start ...
// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name Pieter Lelaona
// @contact.url http://www.swagger.io/support
// @contact.email pieter@rebelworks.co

// @license.name Rebel Private
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://localhost:3000
// @BasePath /v1
func (app API) Start(c *cli.Context) error {
	host := app.Config.GetString("app.host")
	if app.Debug {
		log.Printf("Running on Debug Mode: On at host [%v]", host)
	}

	// gracefull shutdown handler
	valv := valve.New()
	baseCtx := valv.Context()

	// start new app
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"X-SIGNATURE",
			"X-TIMESTAMPT",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)
	if app.Debug {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer)

	RegisterRoutes(r, app.App)

	// handle gracefull shutdown
	srv := http.Server{Addr: host, Handler: chi.ServerBaseContext(baseCtx, r)}
	sng := make(chan os.Signal, 1)
	signal.Notify(sng, os.Interrupt)
	go func() {
		for range sng {
			fmt.Println("shutting down..")
			valv.Shutdown(20 * time.Second)
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()
			srv.Shutdown(ctx)
			select {
			case <-time.After(21 * time.Second):
				fmt.Println("not all connections done")
			case <-ctx.Done():

			}
		}
	}()
	srv.ListenAndServe()

	return nil
}
