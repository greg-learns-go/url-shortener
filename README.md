# URL-shortener

[![CI](https://github.com/greg-learns-go/url-shortener/actions/workflows/ci.yml/badge.svg)](https://github.com/greg-learns-go/url-shortener/actions/workflows/ci.yml)

This is a project to learn GO.

The objective: self-contained app, including a server, that can

- save a new URL, and respond with a short version POST requests
- retrieve a shortened URL on GET, and redirect
- covered with unit tests
- flags specifying domain/port

Bonus points:

- compile an executable for other platforms? (Mac, Windows?)
- deploy the executable in some hosting?

Used concepts:

- ✓ Go basics (slices, loops, pointers etc.)
- ✓ Modules/packages (internal and external)
- ✓ net/http server
- ✓ database/sql
- ✓ func init
- ✓ defer
- ✓ go routines
- ✓ html templates
- ✓ dependency injection (using interfaces)
