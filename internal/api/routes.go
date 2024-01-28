package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, app *AppServer) {
	router.MethodNotAllowedHandler = http.HandlerFunc(app.NotAllowedHandler)
	router.NotFoundHandler = http.HandlerFunc(app.NotFoundHandler)
	router.Methods("GET").Path("/api/health").HandlerFunc(app.Handlers.GetHealth)
	router.Methods("GET").Path("/api/test").HandlerFunc(app.Handlers.GetTest)
	router.Methods("POST").Path("/api/image/create").HandlerFunc(app.Handlers.CreateImage)
	router.Methods("POST").Path("/api/image/create").HandlerFunc(app.Handlers.GetImage)
}
