package main

import (
	"fmt"
	"net"
	"net/http"

	agentmodule "bodyplate.com/cmd/api/module/agent.module"
	postmodule "bodyplate.com/cmd/api/module/post.module"
	"bodyplate.com/cmd/api/route"
	"bodyplate.com/internal/database"
	"google.golang.org/grpc"
)

func main() {
	go func() {
		listen, err := net.Listen("tcp", "127.0.0.1:50052")
		if err != nil {
			fmt.Println("Server gRPC start err: %v", err)
			panic(err)
		}
		fmt.Println("Server gRPC running on: %v", listen.Addr())
		grpcServer := grpc.NewServer()
		agentmodule.Init(grpcServer)

		grpcServer.Serve(listen)
		fmt.Println("INIT SERVER SUCCESS")

	}()
	go func( ) {
	database.DataBaselayer.Init()
	}()
	route.InitRestRoutes()
	postmodule.Init()
	fmt.Println("Server http running on: %v", "6666")
	err := http.ListenAndServe(":6666", route.MuxRoute.Router)
	if err != nil {
		fmt.Println("Server http start err : %v", err)
		panic(err)
	} 
	

	// server gRPC

}
