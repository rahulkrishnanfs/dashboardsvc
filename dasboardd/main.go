package main

import (
	repository "dashboardsvc/cockroachdb/userstorerepository"
	"dashboardsvc/controller"
	"dashboardsvc/router"
	"net/http"
)

func main() {
	useController := controller.NewUserController(repository.NewRepository())
	router := router.InitRoutes(useController)
	http.ListenAndServe(":8000", router)

}
