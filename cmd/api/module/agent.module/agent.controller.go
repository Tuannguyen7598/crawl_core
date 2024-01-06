package agentmodule

import (
	"context"
	"fmt"
	"time"

	agentproto "bodyplate.com/internal/base.grpc/agent.proto"
	"bodyplate.com/internal/database"
	resHttpSchema "bodyplate.com/internal/database/mongo/schema"
)


type AgentController struct {
	AgentServer agentproto.AgentServer
}

func (a *AgentController) SendDataHTTP(ctx context.Context, p *agentproto.DataFromHttpResponse) (*agentproto.ResponseAgent, error) {
	
	data := resHttpSchema.ResponseHttpSchema{
		Sns: p.Sns,
		Url: p.Url,
		Body: p.Body,
		Created_at: time.Now(),
	}
	_,err := database.DataBaselayer.ResHttpRepository.Create(data)
	if err != nil {
		fmt.Println(err)
		return &agentproto.ResponseAgent{Message:"CREATE DATA ERROR",Status: false}, nil
	}
	return &agentproto.ResponseAgent{Message: "CREATE DATA SUCCESS",Status: true}, nil

}
func (a *AgentController) SendDataDOM(ctx context.Context, p *agentproto.DataDOM) (*agentproto.ResponseAgent, error) {
	return &agentproto.ResponseAgent{Message: "Hello World",Status: true}, nil
}

func (a *AgentController) SendActionDOM(ctx context.Context, p *agentproto.ActionDOM) (*agentproto.ResponseAgent, error) {
	fmt.Println(p)

	return &agentproto.ResponseAgent{Message: "Hello World",Status: true}, nil
}