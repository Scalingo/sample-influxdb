package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Scalingo/sample-influxdb/utils"
	"github.com/Scalingo/sample-influxdb/webserver"
	"github.com/Scalingo/sample-influxdb/worker"
)

func main() {
	flagIsWorker := flag.Bool("worker", false, "is a worker")
	flag.Parse()

	var closer utils.Closer
	if *flagIsWorker {
		closer = worker.Start()
	} else {
		closer = webserver.Start()
	}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
	fmt.Println("SIGTERM, time to shutdown")
	if closer != nil {
		_ = closer.Close()
	}
}
