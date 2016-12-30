package influx

import (
	"time"

	"github.com/Scalingo/sample-influxdb/config"
	influx "github.com/influxdata/influxdb/client/v2"
	"gopkg.in/errgo.v1"
)

func Start() (*influx.BatchPoints, error) {
	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  config.InfluxConnectionInformation.Database,
		Precision: "s",
	})
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return &bp, nil
}

func Write(bp *influx.BatchPoints) error {
	client, err := Client()
	if err != nil {
		return errgo.Mask(err)
	}
	defer client.Close()

	err = client.Write(*bp)
	if err != nil {
		return errgo.Mask(err)
	}
	return nil
}

func Add(measurement string, values map[string]interface{}, tags map[string]string, bp *influx.BatchPoints, time time.Time) error {

	pt, err := influx.NewPoint(measurement, tags, values, time)
	if err != nil {
		return errgo.Mask(err)
	}
	(*bp).AddPoint(pt)

	return nil
}
