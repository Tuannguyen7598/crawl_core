package agentmodule

import (
	"context"
	"fmt"
	"time"

	DataDomUseCase "bodyplate.com/domain/dataDom"
	DataHTTPUseCase "bodyplate.com/domain/dataHTTP"
	agentproto "bodyplate.com/internal/base.grpc/agent.proto"
	mongoDatabase "bodyplate.com/internal/repo/mongo"
	"bodyplate.com/internal/until"
)

type AgentController struct {
	AgentServer agentproto.AgentServer
}

func (a *AgentController) SendDataHTTP(ctx context.Context, p *agentproto.DataFromHttpResponse) (*agentproto.ResponseAgent, error) {
	logic := DataHTTPUseCase.DataHttpUseCase{}
	data := mongoDatabase.ResponseHttpSchema{Sns: p.Sns, Url: p.Url, Body: p.Body, Created_at: time.Now()}
	_, err := logic.Create(data)
	if err != nil {
		fmt.Println(err)
		return &agentproto.ResponseAgent{Message: "FALSE", Status: false}, err
	}
	return &agentproto.ResponseAgent{Message: "SUCCESS", Status: true}, nil
}

func (a *AgentController) SendDataDOM(ctx context.Context, p *agentproto.DataDOM) (*agentproto.ResponseAgent, error) {
	logic := DataDomUseCase.DataDomUseCase{}
	until.ConvertEmptyToNil(p)
	_, err := logic.Create(mongoDatabase.DomSchema{Sns: &p.Sns, ActionType: &p.Type, ElementId: &p.Id, ElementName: &p.Name, ElementValue: &p.Value, ElementInnerText: &p.InnerText, Created_at: time.Now()})
	if err != nil {
		fmt.Println(err)
		return &agentproto.ResponseAgent{Message: "FALSE", Status: false}, err
	}
	return &agentproto.ResponseAgent{Message: "SUCCESS", Status: true}, nil
}

func (a *AgentController) SendActionDOM(ctx context.Context, p *agentproto.ActionDOM) (*agentproto.ResponseAgent, error) {
	return &agentproto.ResponseAgent{Message: "NOT_IMPLEMENT", Status: false}, nil
}
