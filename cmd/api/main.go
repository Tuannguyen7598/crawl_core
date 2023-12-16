package main

import (
	"fmt"
	"net"
	"net/http"

	agentmodule "bodyplate.com/cmd/api/module/agent.module"
	postmodule "bodyplate.com/cmd/api/module/post.module"
	"bodyplate.com/cmd/api/route"
	"google.golang.org/grpc"
)

func main() {
	route.InitRestRoutes()
	postmodule.Init()
	err2 := http.ListenAndServe(":5000", route.MuxRoute.Router)
	if err2 != nil {
		fmt.Println("Server http start err : %v", err2)
	}
	// server gRPC
	listen, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println("Server gRPC start err: %v", err)
	}

	grpcServer := grpc.NewServer()
	agentmodule.Init(grpcServer)

	grpcServer.Serve(listen)
	fmt.Println("INIT SERVER SUCCESS")
}
