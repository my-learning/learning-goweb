package main

import (
        "fmt"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
        HostName string = "localhost"
	Address string = ":9090"
)

func mapRoutes() {
	goweb.Map("/", func(c context.Context) error {
		return goweb.Respond.With(c, 200, []byte("Welcome to underground"))
	})

        goweb.Map("/about", func(c context.Context) error {
                return goweb.Respond.With(c, 200, []byte("My name is Keiji Matsuzaki..."))
        })

	goweb.MapStatic("/static", "static")
	goweb.MapStaticFile("/favicon.ico", "static/favicon.ico")
}

func main() {
	mapRoutes()

	log.Print("GoWeb test")
	log.Print(fmt.Sprintf("Starting server: http://%s:%s", HostName, Address))

	s := &http.Server{
		Addr:           Address,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	listener, listenErr := net.Listen("tcp", Address)

	if listenErr != nil {
		log.Fatalf("Cound not listen: %s", listenErr)
	}

	go func() {
		for _ = range c {
			log.Print("Stopping the server...")
			listener.Close()

			log.Print("Tearing down...")
		}
	}()

	log.Fatalf("Error in Serve: %s", s.Serve(listener))

}
