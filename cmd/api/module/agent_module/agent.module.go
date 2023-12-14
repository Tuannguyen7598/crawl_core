package agentmodule

import (
	"fmt"

	BaseModule "bodyplate.com/internal/base_module"
)

func Init() {
	module := BaseModule.BaseModule{
		ModulePrefix:       "post",
		BaseGRPCController: AgentController{},
	}
	fmt.Println(module.ModulePrefix)

}
