package webserver

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/Scalingo/sample-influxdb/config"
	"github.com/Scalingo/sample-influxdb/utils"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Variables struct {
	Hashtag string
}

func Start() utils.Closer {
	m := martini.Classic()
	m.Use(render.Renderer(
		render.Options{
			Directory: "templates",
		},
	))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", Variables{
			Hashtag: config.E["HASHTAG"],
		})
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
