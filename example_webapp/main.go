package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

const (
	Address string = ":9090"
)

// mapRoutes contains lots of examples of how to map things in
// Goweb.  It is in its own function so that test code can call it
// without having to run main().
func mapRoutes() {

	/*
		Add a pre-handler to save the referrer
	*/
	goweb.MapBefore(func(c context.Context) error {

		// add a custom header
		c.HttpResponseWriter().Header().Set("X-Custom-Header", "Goweb")

		return nil
	})

	/*
		Add a post-handler to log something
	*/
	goweb.MapAfter(func(c context.Context) error {
		// TODO: log this
		return nil
	})

	/*
		Map the homepage...
	*/
	goweb.Map("/", func(c context.Context) error {
		return goweb.Respond.With(c, 200, []byte("Welcome to the Goweb example app - see the terminal for instructions."))
	})

	/*
		Map a specific route that will redirect
	*/
	goweb.Map("GET", "people/me", func(c context.Context) error {
		hostname, _ := os.Hostname()
		return goweb.Respond.WithRedirect(c, fmt.Sprintf("/people/%s", hostname))
	})

	/*
		/people (with optional ID)
	*/
	goweb.Map("GET", "people/[id]", func(c context.Context) error {

		if c.PathParams().Has("id") {
			return goweb.API.Respond(c, 200, fmt.Sprintf("Yes, this worked and your ID is %s", c.PathParams().Get("id")), nil)
		} else {
			return goweb.API.Respond(c, 200, "Yes, this worked but you didn't specify an ID", nil)
		}

	})

	/*
		/status-code/xxx
		Where xxx is any HTTP status code.
	*/
	goweb.Map("/status-code/{code}", func(c context.Context) error {

		// get the path value as an integer
		statusCodeInt, statusCodeIntErr := strconv.Atoi(c.PathValue("code"))
		if statusCodeIntErr != nil {
			return goweb.Respond.With(c, http.StatusInternalServerError, []byte("Failed to convert 'code' into a real status code number."))
		}

		// respond with the status
		return goweb.Respond.WithStatusText(c, statusCodeInt)
	})

	// /errortest should throw a system error and be handled by the
	// DefaultHttpHandler().ErrorHandler() Handler.
	goweb.Map("/errortest", func(c context.Context) error {
		return errors.New("This is a test error!")
	})

	/*
		Map a RESTful controller
		(see the ThingsController for all the methods that will get
		 mapped)
	*/
	thingsController := new(ThingsController)
	goweb.MapController(thingsController)

	/*
		Map a handler for if they hit just numbers using the goweb.RegexPath
		function.

		e.g. GET /2468

		NOTE: The goweb.RegexPath is a MatcherFunc, and so comes _after_ the
		      handler.
	*/
	goweb.Map(func(c context.Context) error {
		return goweb.API.RespondWithData(c, "Just a number!")
	}, goweb.RegexPath(`^[0-9]+$`))

	/*
		Map the static-files directory so it's exposed as /static
	*/
	goweb.MapStatic("/static", "static-files")

	/*
		Map the a favicon
	*/
	goweb.MapStaticFile("/favicon.ico", "static-files/favicon.ico")

	/*
		Catch-all handler for everything that we don't understand
	*/
	goweb.Map(func(c context.Context) error {

		// just return a 404 message
		return goweb.API.Respond(c, 404, nil, []string{"File not found"})

	})

}

func main() {

	// map the routes
	mapRoutes()

	/*

	   START OF WEB SERVER CODE

	*/

	log.Print("Goweb 2")
	log.Print("by Mat Ryer and Tyler Bunnell")
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

	log.Println("")
	log.Print("Some things to try in your browser:")
	log.Printf("\t  http://localhost%s", Address)
	log.Printf("\t  http://localhost%s/status-code/404", Address)
	log.Printf("\t  http://localhost%s/people", Address)
	log.Printf("\t  http://localhost%s/people/123", Address)
	log.Printf("\t  http://localhost%s/people/anything", Address)
	log.Printf("\t  http://localhost%s/people/me (will redirect)", Address)
	log.Printf("\t  http://localhost%s/errortest", Address)
	log.Printf("\t  http://localhost%s/things (try RESTful actions)", Address)
	log.Printf("\t  http://localhost%s/123", Address)
	log.Printf("\t  http://localhost%s/static/simple.html", Address)

	log.Println("")
	log.Println("Also try some of these routes:")
	log.Printf("%s", goweb.DefaultHttpHandler())

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

/*
	RESTful example
*/

// Thing is just a thing
type Thing struct {
	Id   string
	Text string
}

// ThingsController is the RESTful MVC controller for Things.
type ThingsController struct {
	// Things holds the things... obviously, you would never do this
	// in the real world - you'd be reading from some kind of datastore.
	Things []*Thing
}

// Before gets called before any other method.
func (r *ThingsController) Before(ctx context.Context) error {

	// set a Things specific header
	ctx.HttpResponseWriter().Header().Set("X-Things-Controller", "true")

	return nil

}

func (r *ThingsController) Create(ctx context.Context) error {

	data, dataErr := ctx.RequestData()

	if dataErr != nil {
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, dataErr.Error())
	}

	dataMap := data.(map[string]interface{})

	thing := new(Thing)
	thing.Id = dataMap["Id"].(string)
	thing.Text = dataMap["Text"].(string)

	r.Things = append(r.Things, thing)

	return goweb.Respond.WithStatus(ctx, http.StatusCreated)

}

func (r *ThingsController) ReadMany(ctx context.Context) error {

	if r.Things == nil {
		r.Things = make([]*Thing, 0)
	}

	return goweb.API.RespondWithData(ctx, r.Things)
}

func (r *ThingsController) Read(id string, ctx context.Context) error {

	for _, thing := range r.Things {
		if thing.Id == id {
			return goweb.API.RespondWithData(ctx, thing)
		}
	}

	return goweb.Respond.WithStatus(ctx, http.StatusNotFound)

}

func (r *ThingsController) DeleteMany(ctx context.Context) error {

	r.Things = make([]*Thing, 0)
	return goweb.Respond.WithOK(ctx)

}

func (r *ThingsController) Delete(id string, ctx context.Context) error {

	newThings := make([]*Thing, 0)

	for _, thing := range r.Things {
		if thing.Id != id {
			newThings = append(newThings, thing)
		}
	}

	r.Things = newThings

	return goweb.Respond.WithOK(ctx)

}
