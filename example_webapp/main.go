package main

import (
	"fmt"
	"github.com/stretchrcom/goweb"
	"github.com/stretchrcom/goweb/context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	Address string = ":9090"
)

func main() {

	/*
		Map some routes
	*/
	goweb.Map("people/me", func(c context.Context) error {
		return goweb.Respond.WithLocation(c, "/people/123")
	})

	goweb.Map("people/[id]", func(c context.Context) error {

		if c.PathParams().Has("id") {
			return goweb.API.Respond(c, 200, fmt.Sprintf("Yes, this worked and your ID is %s", c.PathParams().Get("id")), nil)
		} else {
			return goweb.API.Respond(c, 200, "Yes, this worked but you didn't specify an ID", nil)
		}

		return nil

	})

	/*
		Catch-all handler for everything that we don't understand
	*/
	goweb.Map(func(c context.Context) error {
		goweb.API.Respond(c, 404, nil, []string{"File not found"})
		return nil
	})

	/*

	   START OF WEB SERVER CODE

	*/

	log.Print("Goweb 2")
	log.Print("by Mat Ryer")
	log.Print(" ")
	log.Print("Starting Goweb powered server...")

	// make a http server using the goweb.DefaultHttpHandler()
	s := &http.Server{
		Addr:           Address,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	listener, listenErr := net.Listen("tcp", Address)

	log.Printf("  visit: %s", Address)

	if listenErr != nil {
		log.Fatalf("Could not listen: %s", listenErr)
	}

	go func() {
		for _ = range c {

			// sig is a ^C, handle it

			// stop the HTTP server
			log.Print("Stopping the server...")
			listener.Close()

			/*
			   Tidy up and tear down
			*/
			log.Print("Tearing down...")

			// TODO: tidy code up here

			log.Fatal("Finished - bye bye.  ;-)")

		}
	}()

	// begin the server
	log.Fatalf("Error in Serve: %s", s.Serve(listener))

	/*

	   END OF WEB SERVER CODE

	*/

}
