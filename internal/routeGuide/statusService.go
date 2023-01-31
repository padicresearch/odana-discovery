package routeGuide

import (
	"github.com/padicresearch/odana-discovery/internal/proto"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Status(ctx context.Context, status *proto.NodeStatus) (*emptypb.Empty, error) {
	if err != nil {
		panic(err.Error())
	}
	_, err := db.Collection("status").InsertOne(context.TODO(), status)

	return new(emptypb.Empty), err
}
