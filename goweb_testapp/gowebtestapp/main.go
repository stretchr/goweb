/*
	
	This is a little test app that demonstrates how goweb
	is to be used in a real application.


	INSTALL GOWEB
	=============
	You may need to install 'goweb' in order to import it here.
	To do this navigate to the 'goweb/goweb' directory in a terminal
	and type:
	
		gomake install
	
	This should put a copy of goweb.a into your GOHOME/pkg directory.
		

	BUILD AND RUN THIS TEST APP
	===========================
	To build and run this navigate to the goweb_testapp directory
	in a terminal and type:
	
		gomake; 6l _go_.6; clear; ./6.out; echo ""
		
		
	OR if it says "./6.out: No such file or directory" then
	try the following to see what errors occured:
	
		gomake; 6l _go_.6; ./6.out
	
	
	TRY IT OUT
	==========
	Try out our hosted version on App Engine:
	
		http://gowebtestapp.appspot.com/
		http://gowebtestapp.appspot.com/api.json
		http://gowebtestapp.appspot.com/api/1.json
		http://gowebtestapp.appspot.com/api/1.json?callback=myFunc
		http://gowebtestapp.appspot.com/api/1.json?callback=myFunc&context=123
		
	
	Then in your web browser, visit the following locations:
	
	Controller type:
	
		GET  http://localhost:8080/people/123
		GET  http://localhost:8080/people/456

	Controller func:
	
		GET  http://localhost:8080/animals/Monkey
		GET  http://localhost:8080/animals/Dog

	See different controllers being used for different
	HTTP method:
	
		GET   http://localhost:8080/specific-method/123
		POST  http://localhost:8080/specific-method/123
		
	(Notice the same URL but different HTTP methods)


	RESTful resources
	
		GET			/restful-resource
		GET			/restful-resource/123
		PUT			/restful-resource
		PUT			/restful-resource/123
		POST		/restful-resource
		DELETE	/restful-resource
		DELETE	/restful-resource/123


	RESTful API resources
	
		GET			/api.json
		GET			/api/1.json
		PUT			/api.json
		PUT			/api/1.json
		POST		/api.json
		DELETE	/api.json  		(Forbidden - try it!)
		DELETE	/api/1.json

	Also, see what RespondWithNotFound() looks like:
	
		/api/no-ruch-resource.json
		

	See JSONP in action with these URLs:
	
	  http://localhost:8080/api.json?callback=MyFunc
	  http://localhost:8080/api.json?callback=MyFunc&context=123

*/
package gowebtestapp

import (
	"goweb"
	"fmt"
	"http"
)

/*

	Controller type

*/

// Controller to handle People resources
type PeopleController struct {}
// Handles the /people/{id} requests
func (p *PeopleController) HandleRequest(c *goweb.Context) {
	fmt.Fprintf(c.ResponseWriter, "You are looking for person with ID %s", c.PathParams["id"])
}



/*
	RESTful Controller type
*/

// A resource controller that abides by the RestController
// interface
type MyResourceController struct {}

func (cr *MyResourceController) Create(cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Create new resource")
}
func (cr *MyResourceController) Delete(id string, cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Delete resource %s", id)
}
func (cr *MyResourceController) DeleteMany(cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Delete many resources")
}
func (cr *MyResourceController) Read(id string, cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Read resource %s", id)
}
func (cr *MyResourceController) ReadMany(cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Read many resource")
}
func (cr *MyResourceController) Update(id string, cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Update resource %s", id)
}
func (cr *MyResourceController) UpdateMany(cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Update many resources")
}



/*
	API entity object
*/
type TestEntity struct {
	Id string
	Name string
	Age int
}

/*
	RESTful API Controller type
*/
type MyAPIController struct {}

func (cr *MyAPIController) Create(cx *goweb.Context) {
	cx.RespondWithData(TestEntity{ "1", "Mat", 28 })
}
func (cr *MyAPIController) Delete(id string, cx *goweb.Context) {
	cx.RespondWithOK()
}
func (cr *MyAPIController) DeleteMany(cx *goweb.Context) {
	cx.RespondWithStatus(http.StatusForbidden)
}
func (cr *MyAPIController) Read(id string, cx *goweb.Context) {
	
	if id == "1" {
		cx.RespondWithData(TestEntity{ id, "Mat", 28 })
	} else if id == "2" {
		cx.RespondWithData(TestEntity{ id, "Laurie", 27 })
	} else {
		cx.RespondWithNotFound()
	}
	
}
func (cr *MyAPIController) ReadMany(cx *goweb.Context) {
	cx.RespondWithData([]TestEntity{ TestEntity{ "1", "Mat", 28 }, TestEntity { "2", "Laurie", 27 } })
}
func (cr *MyAPIController) Update(id string, cx *goweb.Context) {
	cx.RespondWithData(TestEntity{ id, "Mat", 28 })
}
func (cr *MyAPIController) UpdateMany(cx *goweb.Context) {
	cx.RespondWithData(TestEntity{ "1", "Mat", 28 })
}


/*
	
	The main function will register the relevant controllers
	and start the web server

*/
func init() {
	
	/*
		Controller type
	*/
	
	// Create a 'people' controller instance (defined above)
	var peopleController *PeopleController = new(PeopleController)
	
	// bind the people controller to the /people/{id} route
	goweb.Map("/people/{id}", peopleController)
	
	
	/*
		ControllerFunc type
	*/
	
	goweb.MapFunc("/animals/{animal_name}", func(c *goweb.Context){
		fmt.Fprintf(c.ResponseWriter, "Your favourite animal is a %s", c.PathParams["animal_name"])
	})
	
	
	/*
		Different GET and POST controllers
	*/
	goweb.MapFunc("/specific-method/{id}", func(c *goweb.Context){
		fmt.Fprintf(c.ResponseWriter, "I will return resource %s.", c.PathParams["id"])
	}, goweb.GetMethod)
	goweb.MapFunc("/specific-method/{id}", func(c *goweb.Context){
		fmt.Fprintf(c.ResponseWriter, "I will create a new resource for %s.", c.PathParams["id"])
	}, goweb.PostMethod)
	
	
	/*
		RESTful resources
	*/
	controller := new(MyResourceController)
	goweb.MapRest("/restful-resource", controller)
	
	
	/*
		API controller
	*/
	apiController := new(MyAPIController)
	goweb.MapRest("/api", apiController)

	goweb.ConfigureDefaultFormatters()
	http.Handle("/", goweb.DefaultHttpHandler)
	
}
