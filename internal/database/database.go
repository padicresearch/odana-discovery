package database

import (
	"context"
	"fmt"
	"github.com/padicresearch/odana-discovery/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() (*mongo.Database, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(util.GetVariableWith("DB_URL")))
	fmt.Println("Db Connected")
	return client.Database("odana"), err
}
