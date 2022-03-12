package cmd

import (
	"tailor/pkg/controllers"
	"tailor/pkg/controllers/auth"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Route() *chi.Mux {
	router := chi.NewRouter()

	router.Use(CORS)
	router.Use(middleware.Recoverer)
	router.Use(ContentTypeMiddleware)
	router.Use(JsonTypeTypeMiddleware)

	router.Route("/products", func(r chi.Router) {
		productGroup := r.Group(nil)
		//	productGroup.Use(JsonTypeTypeMiddleware)
		//	productGroup.Get("/", controllers.GetProductsEndPoint)
		productGroup.Post("/", controllers.CreateEndPoint)
		productGroup.Post("/like/{productId}", controllers.LikeEndPoint)
	})

	router.Route("/auth", func(r chi.Router) {
		authGroup := r.Group(nil)
		authGroup.Post("/login", auth.Login)
	})
	return router
}
