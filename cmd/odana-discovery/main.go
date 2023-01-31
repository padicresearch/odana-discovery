package main

import (
	pb "github.com/padicresearch/odana-discovery/internal/proto"
	"github.com/padicresearch/odana-discovery/internal/routeGuide"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	con, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTelemetryServiceServer(grpcServer, &routeGuide.Server{})

	err = grpcServer.Serve(con)
	if err != nil {
		log.Fatal(err)
	}

}
