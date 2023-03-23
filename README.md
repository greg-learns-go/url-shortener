[![Go](https://github.com/greg-learns-go/url-shortener/actions/workflows/go.yml/badge.svg)](https://github.com/greg-learns-go/url-shortener/actions/workflows/go.yml)

# url-shortener

This is a project to learn GO.

The objective: self contained app, including a server, that can
- save new URL, and respond with a short version on POST requests
- retrieve a shortened URL on GET, and redirect
- covered with unit tests
- flags specifying domain/port

Bonus points:
- compile an executable for other platforms? (Mac, Windows?)
- deploy the executable in some hosting?


Used concepts: 
- [x] Go basics (slices, loops, pointers etc.) 
- [x] Modules/packages (internal and external) 
- [x] net/http server
- [x] database/sql
- [x] func init
- [x] defer
- [ ] go routines
- [ ] html templates
