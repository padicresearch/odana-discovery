package routeGuide

import (
	"fmt"
	"github.com/padicresearch/odana-discovery/internal/proto"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Status(ctx context.Context, status *proto.NodeStatus) (*emptypb.Empty, error) {
	fmt.Println("Node Status")
	return nil, nil
}
