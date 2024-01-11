package postmodule

import (
	"fmt"

	Route "bodyplate.com/cmd/api/route"
	BaseModule "bodyplate.com/internal/base.module"
)

func Init() {
	module := BaseModule.BaseModule{
		ModulePrefix:       "post",
		BaseCRUDController: PostController{},
	}
	fmt.Println("Module", module.ModulePrefix, "init")
	module.InitBaseCRUDRouter(Route.MuxRoute.Router)

}
