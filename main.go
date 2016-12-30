package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Scalingo/sample-influxdb/config"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Closer interface {
	Close() error
}

func main() {
	flagIsWorker := flag.Bool("worker", false, "is a worker")
	flag.Parse()

	var closer Closer
	if *flagIsWorker {
		closer = startWorker()
	} else {
		closer = startServer()
	}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
	fmt.Println("SIGTERM, time to shutdown")
	if closer != nil {
		_ = closer.Close()
	}
}

func startServer() Closer {
	m := martini.Classic()
	m.Use(render.Renderer(
		render.Options{
			Directory: "templates",
		},
	))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	go http.Serve(listener, m)
	log.Println("Listening on 0.0.0.0:" + port)
	return listener
}

func startWorker() Closer {
	conf := oauth1.NewConfig(config.E["TWITTER_CONSUMER_KEY"], config.E["TWITTER_CONSUMER_SECRET"])
	token := oauth1.NewToken(config.E["TWITTER_ACCESS_TOKEN"], config.E["TWITTER_ACCESS_SECRET"])
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := conf.Client(oauth1.NoContext, token)

	// Twitter Client
	client := twitter.NewClient(httpClient)
	_ = client
	demux := twitter.NewSwitchDemux()
	_ = demux
	return nil
}
