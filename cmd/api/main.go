package main

import (
	"fmt"
	"net"
	"net/http"

	agentmodule "bodyplate.com/cmd/api/module/agent.module"
	postmodule "bodyplate.com/cmd/api/module/post.module"
	Route "bodyplate.com/cmd/api/route"

	Database "bodyplate.com/internal/repo"
	Service "bodyplate.com/internal/service"
	"google.golang.org/grpc"
)

func main() {
	go func() {
		listen, err := net.Listen("tcp", "127.0.0.1:50052")
		if err != nil {
			fmt.Printf("Server gRPC start err: %v\n", err)
			panic(err)
		}
		fmt.Printf("Server gRPC running on: %v\n", listen.Addr())
		grpcServer := grpc.NewServer()
		agentmodule.Init(grpcServer)
		grpcServer.Serve(listen)

	}()
	go func( ) {
	Database.DataBaselayer.Init()
	Service.ServiceLayer.Init()
	}()
	Route.InitRestRoutes()
	
	fmt.Printf("Server http running on: %s\n", "6666")
	postmodule.Init()
	err := http.ListenAndServe(":6666", Route.MuxRoute.Router)
	
	if err != nil {
		fmt.Println("Server http start err :", err)
		panic(err)
	} 

}
