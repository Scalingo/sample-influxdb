package config

import (
	"log"
	"net/url"
	"os"

	errgo "gopkg.in/errgo.v1"
)

type InfluxInfo struct {
	Host             string
	User             string
	Password         string
	Database         string
	ConnectionString string
}

var (
	E = map[string]string{
		"INFLUX_URL": "",
		// Twitter credentials to access Twitter API
		"TWITTER_CONSUMER_KEY":    "",
		"TWITTER_CONSUMER_SECRET": "",
		"TWITTER_ACCESS_TOKEN":    "",
		"TWITTER_ACCESS_SECRET":   "",
	}

	InfluxConnectionInformation *InfluxInfo
)

func init() {
	for variable := range E {
		if v := os.Getenv(variable); v != "" {
			E[variable] = v
		}
	}

	var err error
	InfluxConnectionInformation, err = parseConnectionString(E["INFLUX_URL"])
	if err != nil {
		log.Fatalf("Cannot parse the INFLUX_URL connection string: %+v\n", err)
	}
	log.Printf("Will use the following database: %+v\n", InfluxConnectionInformation)
}

func parseConnectionString(con string) (*InfluxInfo, error) {
	url, err := url.Parse(con)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	var password, username string
	if url.User != nil {
		password, _ = url.User.Password()
		username = url.User.Username()
	}

	return &InfluxInfo{
		Host:             url.Scheme + "://" + url.Host,
		User:             username,
		Password:         password,
		Database:         url.Path[1:],
		ConnectionString: con,
	}, nil
}
