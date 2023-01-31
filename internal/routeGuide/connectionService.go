package routeGuide

import (
	"github.com/padicresearch/odana-discovery/internal/proto"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) EstablishConnection(ctx context.Context, connection *proto.Connection) (*emptypb.Empty, error) {
	if err != nil {
		panic(err.Error())
	}

	_, err := db.Collection("node").InsertOne(context.TODO(), connection)

	return new(emptypb.Empty), err
}
