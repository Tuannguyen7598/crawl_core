package route

import (
	"bodyplate.com/cmd/api/midleware"
	"github.com/gorilla/mux"
)

type Mux struct {
	Router *mux.Router
}

var MuxRoute = Mux{
	Router: mux.NewRouter(),
}

func GetRouter() *mux.Router {
	return MuxRoute.Router
}

func InitRestRoutes() {
	MuxRoute.Router.Use(midleware.MiddlewareGlobal)
	MuxRoute.Router.StrictSlash(true)

}
