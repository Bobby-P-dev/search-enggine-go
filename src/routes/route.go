package routes

import (
	"net/http"

	"github.com/Bobby-P-dev/search-enggine-go/src/controllers"
	"github.com/Bobby-P-dev/search-enggine-go/src/middlewares"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(userController controllers.UserController, deviceController controllers.DeviceController) http.Handler {
	r := chi.NewRouter()
	r.Use(middlewares.CorsMiddleware)

	r.Get("/api/v1/users", userController.GetUser)
	r.Post("/api/v1/users/store", userController.CreateUser)

	r.Get("/api/v1/devices", deviceController.GetDevice)
	r.Post("/api/v1/devices/lamp/set", deviceController.SetLamp)

	return r
}
