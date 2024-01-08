package agentmodule

import (
	"fmt"

	BasegRPC "bodyplate.com/internal/base.grpc"
	BaseModule "bodyplate.com/internal/base.module"
	"google.golang.org/grpc"
)

func Init(s *grpc.Server) {
	module := BaseModule.BaseModule{
		ModulePrefix: "agent",
		BaseGRPCController: BasegRPC.BaseGRPCServer{
			AgentServer: &AgentController{},
		},
	}
	fmt.Println("Module", module.ModulePrefix, "init")
	module.InitGRPCRouter(s)
}
