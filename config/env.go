package config

import (
	"log"
	"os"
)

var (
	E = map[string]string{
		// Twitter credentials to access Twitter API
		"TWITTER_CONSUMER_KEY":    "",
		"TWITTER_CONSUMER_SECRET": "",
		"TWITTER_ACCESS_TOKEN":    "",
		"TWITTER_ACCESS_SECRET":   "",
	}
)

func init() {
	for variable := range E {
		if v := os.Getenv(variable); v != "" {
			E[variable] = v
		}
	}

	if E["TWITTER_CONSUER_KEY"] == "" || E["TWITTER_CONSUMER_SECRET"] == "" || E["TWITTER_ACCESS_TOKEN"] == "" || E["TWITTER_ACCESS_SECRET"] == "" {
		log.Fatal("Missing Twitter OAuth1 information")
	}
}
