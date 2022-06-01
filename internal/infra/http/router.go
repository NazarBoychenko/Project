package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"server/internal/infra/http/controllers"
)

func Router(eventController *controllers.EventController) http.Handler {
	router := chi.NewRouter()

	// Health

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes)

		apiRouter.Route("/v1", func(apiRouter chi.Router) {

			apiRouter.Group(func(apiRouter chi.Router) {
				apiRouter.Route("/events", func(apiRouter chi.Router) {
					apiRouter.Get(
						"/",
						eventController.FindAll(),
					)
					apiRouter.Get(
						"/{id}",
						eventController.FindOne(),
					)
					apiRouter.Put(
						"/update",
						eventController.Update(),
					)
					apiRouter.Post(
						"/insert",
						eventController.Insert(),
					)
					apiRouter.Delete(
						"/delete",
						eventController.Delete(),
					)
				})
			})
			apiRouter.Handle("/*", NotFoundBD())
		})
	})

	return router
}
