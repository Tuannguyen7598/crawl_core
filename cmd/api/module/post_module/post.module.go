package Postmodule

import (
	"bodyplate.com/cmd/api/route"
	BaseModule "bodyplate.com/internal/base_module"
)

func Init() {
	module := BaseModule.BaseModule{
		ModulePrefix:       "post",
		BaseCRUDController: PostController{},
	}
	module.InitBaseCRUDRouter(route.MuxRoute.Router)
	module.BaseCRUDRouter.Sub.Path("/id").Methods("POST").HandlerFunc(GetPostByIdPost)

}
