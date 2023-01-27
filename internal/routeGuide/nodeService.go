package routeGuide

import (
	"fmt"
	"github.com/padicresearch/odana-discovery/internal/proto"
	"golang.org/x/net/context"
)

func (s *Server) FindClosestPeers(ctx context.Context, request *proto.FindClosestPeersRequest) (*proto.FindClosestPeersResponse, error) {
	fmt.Println("Find Closet Peers")
	return nil, nil
}
