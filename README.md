# Goweb 2

  * BETA is ready to use

![Alt text](GowebLogoBig.jpg "Goweb 2 - Logo")
  
## Features

  * Drastically improved path matching
  * Cleaner interface for responding (e.g. `goweb.API.RespondWithData`, and `goweb.Respond.WithRedirect`)
  * More control over standard response object for API responses
  * Cleaner RESTful interface design
    * Default OPTIONS implementation that informs clients what methods the controller exposes
  * Much easier to write testable code
  * Better package structure
  * Modular design, making adding new stuff easy
  * Handler mechanism to easily add pre and post handlers to certain requests
  * Uses [stretchrcom/codecs](https://github.com/stretchrcom/codecs) package allowing better support for multiple formats
  * Better error management
  * Performance improvements

## Get started

  * Clone the goweb `v2` branch into `$GOPATH/src/github.com/stretchrcom`.
  * Import the package as usual with `import "github.com/stretchrcom/goweb"` in your code.
  * Open the docs locally by navigating to your `$GOPATH` folder and doing `godoc -http=:6060` in Terminal, then visit `http://localhost:6060/pkg/github.com/stretchrcom/goweb/`.
  * Look at the [example_webapp](https://github.com/stretchrcom/goweb/blob/v2/example_webapp/main.go) project for some ideas of how to get going.
