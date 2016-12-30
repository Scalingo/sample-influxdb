package worker

import (
	"fmt"
	"log"
	"time"

	"github.com/Scalingo/sample-influxdb/config"
	"github.com/Scalingo/sample-influxdb/influx"
	"github.com/Scalingo/sample-influxdb/utils"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type workerCloser struct {
	stream *twitter.Stream
}

func (w workerCloser) Close() error {
	w.stream.Stop()
	return nil
}

func Start() utils.Closer {
	conf := oauth1.NewConfig(config.E["TWITTER_CONSUMER_KEY"], config.E["TWITTER_CONSUMER_SECRET"])
	token := oauth1.NewToken(config.E["TWITTER_ACCESS_TOKEN"], config.E["TWITTER_ACCESS_SECRET"])
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := conf.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		addTweet(tweet.CreatedAt, "tweet")
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		addTweet(dm.CreatedAt, "DM")
	}
	demux.Event = func(event *twitter.Event) {
		addTweet(event.CreatedAt, "event")
	}
	fmt.Println("Starting Stream...")

	filterParams := &twitter.StreamFilterParams{
		Track:         []string{"#osezdirenon"},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}
	closer := workerCloser{
		stream: stream,
	}
	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	return closer
}

func addTweet(createdAt, t string) {
	date := parseTime(createdAt)
	if date == nil {
		return
	}

	bp, err := influx.Start()
	if err != nil {
		log.Printf("Error starting the InfluxDB writer: %+v\n", err)
		return
	}
	err = influx.Add("tweets", map[string]interface{}{"value": 1},
		map[string]string{"type": t}, bp, *date)
	if err != nil {
		log.Printf("Error adding new point to InfluxDB: %+v\n", err)
		return
	}
	err = influx.Write(bp)
	if err != nil {
		log.Printf("Error writing in InfluxDB: %+v\n", err)
		return
	}
	log.Println("Tweet added")
}

func parseTime(date string) *time.Time {
	t, err := time.Parse(time.RubyDate, date)
	if err != nil {
		log.Printf("cannot parse the date: %+v\n", err)
		return nil
	}
	return &t
}
