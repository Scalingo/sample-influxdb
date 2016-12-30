package worker

import (
	"fmt"
	"log"

	"github.com/Scalingo/sample-influxdb/config"
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

	// Twitter Client
	client := twitter.NewClient(httpClient)
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}
	demux.Event = func(event *twitter.Event) {
		fmt.Printf("%#v\n", event)
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
