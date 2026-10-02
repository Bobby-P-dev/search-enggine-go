package routes

import (
	"net/http"

	"github.com/Bobby-P-dev/search-enggine-go/src/controllers"
	"github.com/Bobby-P-dev/search-enggine-go/src/middlewares"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(userController controllers.UserController) http.Handler {
	r := chi.NewRouter()
	r.Use(middlewares.CorsMiddleware)

	r.Get("/api/v1/users", userController.GetUser)
	r.Post("/api/v1/users/store", userController.CreateUser)

	return r
}
