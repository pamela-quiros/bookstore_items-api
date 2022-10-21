package app

import (
	"net/http"

	"github.com/pamela-quiros/bookstore_items-api/src/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

}
