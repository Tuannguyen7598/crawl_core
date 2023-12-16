package agentmodule

import (
	"context"
	"fmt"

	agentproto "bodyplate.com/internal/base.grpc/agent.proto"
)

type Dto struct {
	Name int `json:"name"`
	Age  string
}

type AgentController struct {
	AgentServer agentproto.AgentServer
}

func (a *AgentController) GetFeature(ctx context.Context, p *agentproto.Point) (*agentproto.Feature, error) {
	fmt.Println(p)
	return &agentproto.Feature{Location: &agentproto.Point{Latitude: 1, Longitude: 2}}, nil
}
