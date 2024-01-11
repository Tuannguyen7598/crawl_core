package service

import (
	grpcservice "bodyplate.com/internal/service/grpc.service"
)


type ServiceLayerStruct struct {
	GRPCService grpcservice.GRPCClientService
}

var ServiceLayer = ServiceLayerStruct{}

// Init initializes the ServiceLayerStruct.
//
// No parameters.
// No return types.
func (s *ServiceLayerStruct) Init() {
	s.GRPCService = grpcservice.GRPCClientService{}
	s.GRPCService.Init()
	//
}