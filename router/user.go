package router

import (
	"dashboardsvc/controller"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router, controller *controller.UserController) *mux.Router {
	router.HandleFunc("/signup", controller.Create).Methods("POST")
	router.HandleFunc("/signin", controller.SignIn).Methods("POST")
	return router
}
