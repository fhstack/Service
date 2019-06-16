package main

import (
	"flag"
	"fmt"
	"github.com/l-f-h/Service/generated"
	"github.com/l-f-h/Service/method"
	grpclb "github.com/l-f-h/grpc-lb/etcdv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

var (
	service  = "JudgeService"
	host     = flag.String("host", "127.0.0.1", "Service listening host")
	port     = flag.String("port", "50001", "Service listening port")
	register = flag.String("register", "127.0.0.1:2379", "Etcd address")
)

func main() {
	warn()
	flag.Parse()

	lis, err := net.Listen("tcp", net.JoinHostPort(*host, *port))
	if err != nil {
		log.Fatalf("TCP listen failed: %v", err)
	}
	//With no ttl
	err = grpclb.Register(*register, service, *host, *port, false, -1)
	if err != nil {
		log.Fatalf("Register service: '%s' failed: %v", service, err)
	}
	log.Println("Register successfully!")
	signal()
	server := grpc.NewServer()
	pb_gen.RegisterJudgeServiceServer(server, &method.Server{})
	reflection.Register(server)
	log.Println("Service is ok")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Service failed bind lis on server: %v", err)
	}
}

func signal() {
	ch := make(chan os.Signal, 1)
	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		grpclb.UnRegister()
		os.Exit(1)
	}()
}

func warn() {
	fmt.Println("Must Run the JudgeService under the ROOT user!!!!!")
	fmt.Println("Must Run the JudgeService under the ROOT user!!!!!")
	fmt.Println("Must Run the JudgeService under the ROOT user!!!!!")
}
