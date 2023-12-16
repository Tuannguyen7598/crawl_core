package BasegRPC

import (
	agentproto "bodyplate.com/internal/base.grpc/agent.proto"
)

type BaseServerGRPC struct {
	AgentServer agentproto.AgentServer
}
