package Baserestapi

import "net/http"

type IBaseRestCRUDConTroller interface {
	HandleGetRoute(w http.ResponseWriter, r *http.Request)
	HandlePostRoute(w http.ResponseWriter, r *http.Request)
	HandlePutRoute(w http.ResponseWriter, r *http.Request)
	HandleDeleteRoute(w http.ResponseWriter, r *http.Request)
}
