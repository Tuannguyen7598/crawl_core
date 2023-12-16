package postmodule

import (
	"fmt"

	"bodyplate.com/cmd/api/route"
	BaseModule "bodyplate.com/internal/base.module"
)

func Init() {
	module := BaseModule.BaseModule{
		ModulePrefix:       "post",
		BaseCRUDController: PostController{},
	}
	fmt.Println("Module", module.ModulePrefix, "init")
	module.InitBaseCRUDRouter(route.MuxRoute.Router)
	module.BaseCRUDRouter.Sub.Path("/id").Methods("POST").HandlerFunc(GetPostByIdPost)

}
