package app

import (
	"github.com/ArturoAguilar1/b_items-api/controllers"
	"net/http"
)

func mapUrls(){
	//router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}