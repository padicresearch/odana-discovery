package routeGuide

import (
	"github.com/padicresearch/odana-discovery/internal/database"
	"github.com/padicresearch/odana-discovery/internal/proto"
)

var db, err = database.ConnectToMongoDB()

type Server struct {
	proto.UnimplementedTelemetryServiceServer
}
