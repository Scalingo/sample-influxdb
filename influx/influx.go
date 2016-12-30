package influx

import (
	"log"

	"github.com/Scalingo/sample-influxdb/config"
	influx "github.com/influxdata/influxdb/client/v2"

	"gopkg.in/errgo.v1"
)

func Client() (influx.Client, error) {
	client, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     config.InfluxConnectionInformation.Host,
		Username: config.InfluxConnectionInformation.User,
		Password: config.InfluxConnectionInformation.Password,
	})

	if err != nil {
		return nil, errgo.Mask(err)
	}

	return client, err
}

func CreateDatabase() {
	queryString := "CREATE DATABASE tweets"
	_, _, err := executeQuery(queryString)
	if err != nil {
		log.Fatalf("Cannot create the database: %+v\n", err)
	}
}
