package api

import (
	"sub/internals/app/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes(orderHandler *handlers.OrdersHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/orders/{id}", orderHandler.FindByID).Methods("GET")
	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()
	return r
}
