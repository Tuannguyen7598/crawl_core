package route

import (
	"bodyplate.com/cmd/api/midleware"
	routeguide "bodyplate.com/internal/base_grpc/proto"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
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

func InitGrpcServer(s *grpc.Server) {
	routeguide.RegisterRouteGuideServer(s, &Server{})
}
