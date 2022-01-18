package cmd

import (
	"tailor/pkg/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Route() *chi.Mux {
	router := chi.NewRouter()
	router.Use(CORS)
	router.Use(middleware.Recoverer)
	router.Use(ContentTypeMiddleware)

	router.Route("/products", func(r chi.Router) {
		productGroup := r.Group(nil)
		productGroup.Use(JsonTypeTypeMiddleware)
		//	productGroup.Get("/", controllers.GetProductsEndPoint)
		productGroup.Post("/", controllers.CreateEndPoint)
		productGroup.Post("/like/{productId}", controllers.LikeEndPoint)
	})
	return router
}
