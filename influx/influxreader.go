package influx

import (
	"errors"
	"time"

	"github.com/Scalingo/sample-influxdb/config"
	influx "github.com/influxdata/influxdb/client/v2"
	"gopkg.in/errgo.v1"
)

type InfluxValue struct {
	Time  time.Time   `json:"time"`
	Value interface{} `json:"value"`
}

// For the last values we took the two last values. The controller will only take the second one. This prevent null value or incomplete aggregation.

func LastMinutesTweets() ([]InfluxValue, bool, error) {
	queryString := "SELECT count(\"value\") FROM \"tweets\""
	queryString += " WHERE hashtag = '" + config.E["HASHTAG"] + "'"
	queryString += " AND   time >= now() - " + config.E["LAST_MINUTES"] + "m"
	queryString += " GROUP BY time(" + config.E["LAST_MINUTES"] + "m) fill(none) ORDER BY time DESC LIMIT " + config.E["LAST_MINUTES"]

	return executeQuery(queryString)
}

func executeQuery(queryString string) ([]InfluxValue, bool, error) {
	client, err := Client()
	if err != nil {
		return nil, true, errgo.Mask(err)
	}
	defer client.Close()

	response, err := client.Query(influx.Query{
		Command:  queryString,
		Database: config.InfluxConnectionInformation.Database,
	})

	if err != nil {
		return nil, true, errgo.Mask(err)
	}
	if response.Error() != nil {
		return nil, true, errgo.Mask(response.Error())
	}

	results, found, err := convertResults(response)

	if err != nil {
		return nil, true, errgo.Mask(err)
	}

	return results, found, nil
}

func convertResults(response *influx.Response) ([]InfluxValue, bool, error) {
	if len(response.Results) < 1 || len(response.Results[0].Series) < 1 {
		return nil, false, nil
	}

	results := make([]InfluxValue, 0, len(response.Results[0].Series[0].Values))

	for _, row := range response.Results[0].Series[0].Values {
		time, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			return nil, true, errors.New("Internal error : Invalid time")
		}
		val := row[1]
		results = append(results, InfluxValue{
			Time:  time,
			Value: val,
		})
	}
	return results, true, nil
}
