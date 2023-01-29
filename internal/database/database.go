package database

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/padicresearch/odana-discovery/internal/util"
)

func DatabaseConfig() api.WriteAPIBlocking {
	bucket := util.GetVariableWith("BUCKET")
	org := util.GetVariableWith("ORG")
	token := util.GetVariableWith("INFLUX_ODANA_TOKEN")
	url := util.GetVariableWith("INFLUX_URL")
	client := influxdb2.NewClient(url, token)
	return client.WriteAPIBlocking(org, bucket)
}
