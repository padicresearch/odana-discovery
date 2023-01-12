package database

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func DatabaseConfig() api.WriteAPIBlocking {
	bucket := "example-bucket"
	org := "example-org"
	token := "example-token"
	// Store the URL of your InfluxDB instance
	url := "localhost:8086"
	client := influxdb2.NewClient(url, token)
	return client.WriteAPIBlocking(org, bucket)
}
