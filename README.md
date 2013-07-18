# Goweb

NOTE: We have changed our repo name from `stretchrcom` to just `stretchr`... please update all your paths (find and replace on `github.com/stretchr` should work great).  We apologise for any inconvenience caused.

A lightweight RESTful web framework for Go.

![Goweb A lightweight RESTful web framework for Go.](GowebLogoBig.jpg "Goweb 2 - Logo")

  * If you'd like to chat about Goweb, please feel free to join our [HipChat Goweb Channel](http://www.hipchat.com/gXWgwTtX2)
  * For examples and usage, please read the [Goweb API Documentation](http://godoc.org/github.com/stretchr/goweb)

## Get started

  * To install, run `go get github.com/stretchr/goweb`
  * Import the package as usual with `import "github.com/stretchr/goweb"` in your code.
  * Look at the [example_webapp](https://github.com/stretchr/goweb/blob/master/example_webapp/main.go) project for some ideas of how to get going
  * Read the [Goweb API Documentation](http://godoc.org/github.com/stretchr/goweb)
  * To update to the latest version of goweb, just run `go get -u github.com/stretchr/goweb`

## Features

  * Drastically improved path matching
  * Cleaner interface for responding (e.g. `goweb.API.RespondWithData`, and `goweb.Respond.WithRedirect`)
  * More control over standard response object for API responses
  * Cleaner RESTful interface design
    * Default OPTIONS implementation that informs clients what methods the controller exposes
  * Easily publish static files as well as code driven output
  * Much easier to write testable code
  * Better package structure
  * Modular design, making adding new stuff easy
  * Handler mechanism to easily add pre and post handlers to certain requests
  * Uses [stretchr/codecs](https://github.com/stretchr/codecs) package allowing better support for multiple formats
  * Easily match paths using Regex instead
  * Better error management
  * Performance improvements

## Requirements

  * Goweb runs on Go 1.1

------

Contributing
============

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include steps to reproduce the issue so we can see it on our end also!


Licence
=======
Copyright (c) 2012 - 2013 Mat Ryer and Tyler Bunnell

Please consider promoting this project if you find it useful.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
