
package router

import (
    "net/http"
    "gocool/internal/controller"
    "gocool/internal/middleware"
    "github.com/gorilla/mux"
)

// SetupRoutes configures the routes for the application
func SetupRoutes(userController *controller.UserController) *mux.Router {
    r := mux.NewRouter()

    // User routes
    r.HandleFunc("/user/register", userController.RegisterUser).Methods("POST")
    r.HandleFunc("/user/details", userController.GetUserDetails).Methods("GET")

    // Apply middleware
    r.Use(middleware.TracerMiddleware)
    r.Use(middleware.AuthMiddleware)

    return r
}
