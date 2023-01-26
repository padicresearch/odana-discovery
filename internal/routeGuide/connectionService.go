package routeGuide

import (
	"github.com/padicresearch/odana-discovery/internal/proto"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) EstablishConnection(ctx context.Context, connection *proto.Connection) (*emptypb.Empty, error) {
	log.Printf("Data recieved %s", connection)
	return nil, nil
}
