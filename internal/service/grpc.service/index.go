package grpcService

import (
	"fmt"

	"bodyplate.com/config"
	agentproto "bodyplate.com/internal/base.grpc/agent.proto"
	"google.golang.org/grpc"
)

type GRPCClientService struct {
	AgentClient agentproto.AgentClient
}


var client = GRPCClientService{}
func (g * GRPCClientService) Init() {
	// Configure the target URL for the gRPC client to call
	targetURL := config.GetConfigAvaiable().Uri_agent // Replace with your actual gRPC server URL
	conn, err := grpc.Dial(targetURL, grpc.WithInsecure()) // Use grpc.WithTransportCredentials for secure connections
	if err != nil {
		panic(fmt.Errorf("failed to connect to gRPC server: %w", err))
	}
	fmt.Println("Service gRPC init")
	g.AgentClient = agentproto.NewAgentClient(conn)
}
