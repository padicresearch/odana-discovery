package main

import (
	"github.com/padicresearch/odana-discovery/internal/config"
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

	grpcServer := grpc.NewServer(grpc.StatsHandler(&config.Handler{}))
	pb.RegisterTelemetryServiceServer(grpcServer, &routeGuide.Server{})

	err = grpcServer.Serve(con)
	if err != nil {
		log.Fatal(err)
	}

}
