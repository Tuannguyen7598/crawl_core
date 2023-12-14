package main

import (
	"fmt"
	"net"
	"net/http"

	Postmodule "bodyplate.com/cmd/api/module/post_module"
	"bodyplate.com/cmd/api/route"
	"google.golang.org/grpc"
)

func main() {
	route.InitRestRoutes()
	Postmodule.Init()
	fmt.Println("Server start at port 5000")
	err2 := http.ListenAndServe(":5000", route.MuxRoute.Router)
	if err2 != nil {
		fmt.Println("Server start err : %v", err2)
	}
	// server gRPC
	listen, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println("False to inti listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	route.InitGrpcServer(grpcServer)

	grpcServer.Serve(listen)
	fmt.Println("INIT SERVER SUCCESS")
}
