package BaseModule

import (
	"fmt"

	BasegRPC "bodyplate.com/internal/base_grpc"
	Baserestapi "bodyplate.com/internal/base_rest_api"
	"github.com/gorilla/mux"
)

type BaseModule struct {
	ModulePrefix       string
	BaseCRUDRouter     Baserestapi.BaseRestCRUDRouter
	BaseCRUDController Baserestapi.IBaseRestCRUDConTroller
	BaseGRPCController BasegRPC.BaseServerGRPC
}

func (b *BaseModule) InitBaseCRUDRouter(router *mux.Router) {
	path := fmt.Sprint("/", b.ModulePrefix)
	subRouter := router.PathPrefix(path).Subrouter()
	b.BaseCRUDRouter.Get = router.Path(path).HandlerFunc(b.BaseCRUDController.HandleGetRoute).Methods("GET")
	b.BaseCRUDRouter.Post = router.Path(path).HandlerFunc(b.BaseCRUDController.HandlePostRoute).Methods("POST")
	b.BaseCRUDRouter.Put = router.Path(path).HandlerFunc(b.BaseCRUDController.HandlePutRoute).Methods("PUT")
	b.BaseCRUDRouter.Put = router.Path(path).HandlerFunc(b.BaseCRUDController.HandleDeleteRoute).Methods("DELETE")
	b.BaseCRUDRouter.Sub = subRouter
}

func (b *BaseModule) InitGRPCRouter() {

}
