package BaseRestApi

import "github.com/gorilla/mux"

type BaseRestCRUDRouter struct {
	Sub    *mux.Router
	Get    *mux.Route
	Post   *mux.Route
	Put    *mux.Route
	Delete *mux.Route
}
