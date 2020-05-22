package main

import (
	"net/http"

	"github.com/dudleycodes/golang-microservice-structure/app/storefront-api/middleware"
	"github.com/dudleycodes/golang-microservice-structure/app/storefront-api/routes"
	"github.com/dudleycodes/golang-microservice-structure/app/storefront-api/routes/makes"
	"github.com/dudleycodes/golang-microservice-structure/app/storefront-api/routes/makes/models"
	"github.com/dudleycodes/golang-microservice-structure/app/storefront-api/webserver"
	"github.com/gorilla/mux"
)

// BindRoutes builds the HTTP pipeline
func BindRoutes(srv webserver.Server, r *mux.Router) {
	r.HandleFunc("/ping", routes.Ping(srv)).Methods(http.MethodGet)

	r.Use(middleware.Authentication(srv))

	r.HandleFunc("`/makes/{makeID}", makes.Get(srv)).Methods(http.MethodGet)
	r.HandleFunc("`/makes/{makeID}/models/{modelID}", models.Get(srv)).Methods(http.MethodGet)
}
