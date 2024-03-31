package main

import (
	"fmt"
	"grpc_calc/internal/handlers/grpc_controller"
	"net"

	"google.golang.org/grpc"
)

func main() {
	gRPCServer := grpc.NewServer()
	grpc_controller.Register(gRPCServer)

	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic("cant listen socket")
	}

	err = gRPCServer.Serve(l)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic("cant serve")
	}
}
