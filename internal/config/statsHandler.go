package config

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/stats"
	"log"
)

type Handler struct {
}

func (h *Handler) TagRPC(context.Context, *stats.RPCTagInfo) context.Context {
	log.Println("Connection Registered")
	return context.Background()
}

func (h *Handler) HandleRPC(context.Context, stats.RPCStats) {
	log.Println("Handle connections")
}

func (h *Handler) TagConn(context.Context, *stats.ConnTagInfo) context.Context {
	log.Println("Tag Connection")
	return context.Background()
}

func (h *Handler) HandleConn(ctx context.Context, status stats.ConnStats) {
	switch status.(type) {
	case *stats.ConnEnd:
		log.Println("Client Has Disconnected")
		break
	}
}
