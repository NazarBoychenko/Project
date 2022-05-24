package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"server/internal/infra/http/controllers"
)

func Router(eventController *controllers.EventController, web *UseWeb) http.Handler {
	router := chi.NewRouter()

	// Health
	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes)

		apiRouter.Route("/v1", func(apiRouter chi.Router) {

			apiRouter.Group(func(apiRouter chi.Router) {
				AddEventRoutes(&apiRouter, eventController)

				apiRouter.Handle("/*", NotFoundBD())
			})
			apiRouter.Handle("/*", NotFoundBD())
		})
	})

	router.Group(func(updateRouter chi.Router) {
		updateRouter.Use(middleware.RedirectSlashes)

		updateRouter.Route("/updateInsert", func(updateRouter chi.Router) {
			updateRouter.Get("/", web.UpdateInsert())
			updateRouter.Post("/", web.UpdateInsert())

			updateRouter.Handle("/*", NotFoundBD())
		})
	})

	return router
}

func AddEventRoutes(router *chi.Router, eventController *controllers.EventController) {
	(*router).Route("/events", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			eventController.FindAll(),
		)
		apiRouter.Get(
			"/{id}",
			eventController.FindOne(),
		)
	})
}
