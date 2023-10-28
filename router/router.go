package router

import (
	"dashboardsvc/controller"

	"github.com/gorilla/mux"
)

func InitRoutes(controller *controller.UserController) *mux.Router {
	router := mux.NewRouter()
	router = UserRoutes(router, controller)
	return router
}
