package main

import (
	"github.com/l-f-h/Service/conf"
	"github.com/l-f-h/Service/generated"
	"github.com/l-f-h/Service/method"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", conf.RpcPort)
	if err != nil {
		log.Fatalf("Service Judge failed listen port %d: %v", conf.RpcPort, err)
	}
	server := grpc.NewServer()
	pb_gen.RegisterJudgeServiceServer(server, &method.Server{})
	reflection.Register(server)
	log.Println("Service is ok")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Service failed bind lis on server: %v", err)
	}
}
